package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"

	// _ "expvar"
	_ "net/http/pprof"

	"github.com/Wessie/hanyuu/streamer"
)

func main() {
	flag.StringVar(&streamer.DefaultConfig.TemplateDir, "templatedir",
		streamer.DefaultConfig.TemplateDir, "directory containing html templates")
	configPath := flag.String("conf", "hanyuu.toml",
		"filepath to configuration file")
	graceful := flag.Bool("graceful", false, "internal use only")

	flag.Parse()

	s, err := streamer.NewState(*configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("gracefulled")

	// we're expecting files from an old process
	if *graceful {
		file := os.NewFile(3, "listener")
		l, err := net.FileListener(file)
		if err != nil {
			// file.Close()
			fmt.Println("graceful: new: failed to inherit socket:", err)
		}
		if err = file.Close(); err != nil {
			fmt.Println("graceful: new: failed closing inherited socket:", err)
		}

		file = os.NewFile(4, "icecast")
		conn, err := net.FileConn(file)
		if err != nil {
			// file.Close()
			fmt.Println("graceful: new: failed to inherit conn:", err)
		}
		if err = file.Close(); err != nil {
			fmt.Println("graceful: new: failed closing inherited conn:", err)
		}

		s.GracefulSetup(l, conn)

		// only start the streamer right away if we're restarting ourselves
		s.StartStreamer()
	}

	fmt.Println(s.Conf())
	go func() {
		err := streamer.ListenAndServe(s)
		if err != nil {
			fmt.Println("http: serve error:", err)
		}
	}()

	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("received signal, exiting process")
	s.Shutdown()
}