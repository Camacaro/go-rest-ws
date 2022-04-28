package models

type WebSocketMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"` // interface{} es una interfaz que puede ser cualquier tipo de dato, es como un any en typescript
}
