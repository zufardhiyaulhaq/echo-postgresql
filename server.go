package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/evio"
	"github.com/zufardhiyaulhaq/echo-postgresql/pkg/settings"
	"github.com/zufardhiyaulhaq/echo-postgresql/pkg/types"

	postgresql_client "github.com/zufardhiyaulhaq/echo-postgresql/pkg/postgresql"
)

type Server struct {
	settings settings.Settings
	client   postgresql_client.Interface
}

func NewServer(settings settings.Settings, client postgresql_client.Interface) Server {
	return Server{
		settings: settings,
		client:   client,
	}
}

func (e Server) ServeEcho() {
	var events evio.Events

	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		id := uuid.NewString()
		value := string(in)

		echo := types.Echo{
			ID:   id,
			Echo: value,
		}

		err := e.client.WriteEcho(&echo)
		if err != nil {
			out = []byte(err.Error())
			return
		}

		read, err := e.client.GetEcho(id)
		if err != nil {
			out = []byte(err.Error())
			return
		}

		out = []byte(read.ID + ":" + read.Echo)

		return
	}

	if err := evio.Serve(events, "tcp://0.0.0.0:"+e.settings.EchoPort); err != nil {
		log.Fatal().Err(err)
	}
}

func (e Server) ServeHTTP() {
	handler := NewHandler(e.settings, e.client)

	r := mux.NewRouter()

	r.HandleFunc("/postgresql/{key}", handler.Handle)
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello!"))
	})
	r.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello!"))
	})

	err := http.ListenAndServe(":"+e.settings.HTTPPort, r)
	if err != nil {
		log.Fatal().Err(err)
	}
}

type Handler struct {
	settings settings.Settings
	client   postgresql_client.Interface
}

func NewHandler(settings settings.Settings, client postgresql_client.Interface) Handler {
	return Handler{
		settings: settings,
		client:   client,
	}
}

func (h Handler) Handle(w http.ResponseWriter, req *http.Request) {
	id := uuid.NewString()
	value := mux.Vars(req)["key"]

	echo := types.Echo{
		ID:   id,
		Echo: value,
	}

	err := h.client.WriteEcho(&echo)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
		return
	}

	read, err := h.client.GetEcho(id)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(read.ID + ":" + read.Echo))
}
