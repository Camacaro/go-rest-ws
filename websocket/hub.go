package websocket

import (
	"encoding/json"
	"fmt"
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

// Ejecutar el hub
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register: // Un cliente se acaba de registrar
			h.onConnect(client)
		case client := <-h.unregister: // Un cliente se acaba de desregistrar
			h.onDisconnect(client)
		}
	}
}

func (h *Hub) onConnect(client *Client) {
	fmt.Println("hub.go: Hub.onConnect: Cliente conectado", client.socket.RemoteAddr())

	h.mutex.Lock()
	defer h.mutex.Unlock()

	client.id = client.socket.RemoteAddr().String()
	h.clients = append(h.clients, client)
}

func (h *Hub) onDisconnect(client *Client) {
	fmt.Println("hub.go: Hub.onDisconnect: Cliente desconectado", client.socket.RemoteAddr())

	// Cerrar la conexion
	client.socket.Close()

	h.mutex.Lock()
	defer h.mutex.Unlock()

	// Eliminar el cliente del array
	// for i, c := range h.clients {
	// 	if c.id == client.id {
	// 		h.clients = append(h.clients[:i], h.clients[i+1:]...)
	// 		break
	// 	}
	// }

	i := -1
	for j, c := range h.clients {
		if c.id == client.id {
			i = j
			break
		}
	}

	copy(h.clients[i:], h.clients[i+1:])
	h.clients[len(h.clients)-1] = nil
	h.clients = h.clients[:len(h.clients)-1]
}

// TRansmitir los mensajes a todos los clientes
// ignore nos sirve para ignorar el cliente que envia el mensaje
func (h *Hub) Broadcast(message interface{}, ignore *Client) {
	// Serializar la data
	data, _ := json.Marshal(message)

	for _, client := range h.clients {
		if client != ignore {
			client.outbound <- data // Enviar el mensaje al canal
		}
	}
}
