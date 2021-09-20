package main

import (
	"github.com/google/uuid"
	"github.com/tidwall/evio"
	"github.com/zufardhiyaulhaq/echo-postgresql/pkg/settings"
	"github.com/zufardhiyaulhaq/echo-postgresql/pkg/types"

	postgresql_client "github.com/zufardhiyaulhaq/echo-postgresql/pkg/postgresql"
)

func main() {
	var events evio.Events

	settings, err := settings.NewSettings()
	if err != nil {
		panic(err.Error())
	}

	store := postgresql_client.New(settings)

	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		id := uuid.NewString()
		value := string(in)

		echo := types.Echo{
			ID:   id,
			Echo: value,
		}

		err := store.WriteEcho(&echo)
		if err != nil {
			out = []byte(err.Error())
			return
		}

		readEcho, err := store.GetEcho(id)
		if err != nil {
			out = []byte(err.Error())
			return
		}

		out = []byte(readEcho.Echo)

		return
	}

	if err := evio.Serve(events, "tcp://0.0.0.0:"+settings.PostgresqlEventPort); err != nil {
		panic(err.Error())
	}

}
