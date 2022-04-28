package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Centralizar todos los clientes y distribuirlos a los canales
var upgrader = websocket.Upgrader{
	/*
		Restringir el acceso para ciertos clientes, pero no es necesario para este proyecto
		Por lo cual todos los clientes se podran conectar
	*/
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Hub struct {
	clients    []*Client
	register   chan *Client // Canal para registrar clientes
	unregister chan *Client // Canal para desregistrar clientes
	mutex      *sync.Mutex  // Evitar condiciones de carreras en el canal
}

func NewHub() *Hub {
	return &Hub{
		clients:    make([]*Client, 0),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}

func (h *Hub) HandleWebScoket(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln("hub.go: Hub.HandleWebScoket: ", err)
		http.Error(w, "could not open websocket connection", http.StatusBadRequest)
		return
	}

	client := NewClient(h, socket)
	h.register <- client // Registra el cliente en el canal de registro

	/*
		Crear una go routine para empezar a escribir los diferentes mensajes
		para este websockets
	*/
	go client.Write()
}
