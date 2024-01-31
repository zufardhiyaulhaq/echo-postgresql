package main

import (
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/zufardhiyaulhaq/echo-postgresql/pkg/settings"

	postgresql_client "github.com/zufardhiyaulhaq/echo-postgresql/pkg/postgresql"
)

func main() {
	settings, err := settings.NewSettings()
	if err != nil {
		panic(err.Error())
	}

	log.Info().Msg("creating postgresql client")
	store := postgresql_client.New(settings)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	log.Info().Msg("starting server")
	server := NewServer(settings, store)

	go func() {
		log.Info().Msg("starting HTTP server")
		server.ServeHTTP()
		wg.Done()
	}()

	go func() {
		log.Info().Msg("starting echo server")
		server.ServeEcho()
		wg.Done()
	}()

	wg.Wait()
}
