package server

import (
	"context"
	"errors"
	"go-rest-ws/database"
	"go-rest-ws/repository"
	"go-rest-ws/websocket"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
	Hub() *websocket.Hub // WebSocket, crearemos nuestros propios hubs (clieente)
}

/*
	Manejar diferentes servers
*/
type Broker struct {
	config *Config
	router *mux.Router
	hub    *websocket.Hub
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Hub() *websocket.Hub {
	return b.hub
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)

	repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	go b.hub.Run()
	repository.SetRepository(repo)

	log.Println("Server listening on port", b.Config().Port)
	if err := http.ListenAndServe(b.Config().Port, b.router); err != nil {
		log.Fatal(err)
	}
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("JWT Secret is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("Database Url is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		hub:    websocket.NewHub(),
	}

	return broker, nil
}
