package jobs

import (
	"context"

	"github.com/R-a-dio/valkyrie/config"
	"github.com/R-a-dio/valkyrie/storage"
)

// ExecuteListenerLog fetches the listener count from the manager and inserts a line into
// the listenlog table.
func ExecuteListenerLog(ctx context.Context, cfg config.Config) error {
	store, err := storage.Open(cfg)
	if err != nil {
		return err
	}

	m := cfg.Conf().Manager.Client()

	status, err := m.Status(ctx)
	if err != nil {
		return err
	}

	return store.User(ctx).RecordListeners(status.Listeners, status.User)
}
