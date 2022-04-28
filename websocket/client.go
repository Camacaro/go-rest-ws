package websocket

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub      *Hub
	id       string
	socket   *websocket.Conn
	outbound chan []byte // Enviar mensajes
}

func NewClient(hub *Hub, socket *websocket.Conn) *Client {
	return &Client{
		hub:      hub,
		socket:   socket,
		outbound: make(chan []byte),
	}
}

func (c *Client) Write() {
	/*
		Escribir los mensajes en el canal de escritura
	*/
	for {
		select {
		case message, ok := <-c.outbound: // Si recibe un mensaje, voy a querer transmitir esa data
			// I fue diferente de ok hubo un problema
			if !ok {
				// Cerrar la conexion
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			fmt.Println("client.go: Client.Write: Enviando mensaje", string(message))

			// Enviar el mensaje - Escribit el mensaje que estoy recibiendo
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
