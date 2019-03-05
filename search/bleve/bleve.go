package bleve

import (
	"context"
	"encoding/json"

	radio "github.com/R-a-dio/valkyrie"
	"github.com/R-a-dio/valkyrie/config"
	"github.com/R-a-dio/valkyrie/errors"
	"github.com/R-a-dio/valkyrie/search"
	"github.com/R-a-dio/valkyrie/storage"

	"github.com/blevesearch/bleve"
)

func init() {
	search.Register("memory", Open)
}

func Open(cfg config.Config) (radio.SearchService, error) {
	const op errors.Op = "bleve/Open"

	mapping := bleve.NewIndexMapping()
	index, err := bleve.NewMemOnly(mapping)
	if err != nil {
		return nil, errors.E(op, err)
	}

	ss := SearchService{index}

	// since we're an in-memory index, we need to reindex everything
	store, err := storage.Open(cfg)
	if err != nil {
		return nil, errors.E(op, err)
	}

	songs, err := store.Track(context.Background()).All()
	if err != nil {
		return nil, errors.E(op, err)
	}

	err = ss.Update(context.Background(), songs...)
	if err != nil {
		return nil, errors.E(op, err)
	}

	return ss, nil
}

type SearchService struct {
	index bleve.Index
}

func (s SearchService) Search(ctx context.Context, query string, limit int, offset int) ([]radio.Song, error) {
	const op errors.Op = "bleve/SearchService.Search"

	match := bleve.NewMatchQuery(query)

	request := bleve.NewSearchRequestOptions(match, limit, offset, false)
	request.SortBy([]string{"-priority"})
	request.Fields = []string{"*"}

	result, err := s.index.SearchInContext(ctx, request)
	if err != nil {
		return nil, errors.E(op, err)
	}

	if result.Hits == nil || len(result.Hits) == 0 {
		return []radio.Song{}, errors.E(op, errors.SearchNoResults, errors.Info(query))
	}

	songs := make([]radio.Song, len(result.Hits))
	for i, hit := range result.Hits {
		b, err := json.Marshal(hit.Fields)
		if err != nil {
			return nil, errors.E(op, err)
		}

		err = json.Unmarshal(b, &songs[i])
		if err != nil {
			return nil, errors.E(op, err)
		}
		songs[i].FillMetadata()
	}

	return songs, nil
}

func (s SearchService) Update(ctx context.Context, songs ...radio.Song) error {
	const op errors.Op = "bleve/SearchService.Update"

	if len(songs) == 1 {
		song := songs[0]
		if !song.HasTrack() {
			return errors.E(op, errors.SongWithoutTrack, song)
		}

		err := s.index.Index(song.TrackID.String(), song)
		if err != nil {
			return errors.E(op, err)
		}
		return nil
	}

	songs = songs[:1000]

	batch := s.index.NewBatch()
	for _, song := range songs {
		if !song.HasTrack() {
			return errors.E(op, errors.SongWithoutTrack, song)
		}

		err := batch.Index(song.TrackID.String(), song)
		if err != nil {
			return errors.E(op, err)
		}
	}

	err := s.index.Batch(batch)
	if err != nil {
		return errors.E(op, err)
	}
	return nil
}

func (s SearchService) Delete(ctx context.Context, songs ...radio.Song) error {
	const op errors.Op = "bleve/SearchService.Delete"

	for _, song := range songs {
		if !song.HasTrack() {
			return errors.E(op, errors.SongWithoutTrack, song)
		}

		err := s.index.Delete(song.TrackID.String())
		if err != nil {
			return errors.E(op, err)
		}
	}
	return nil
}
