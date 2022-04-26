package handlers

import (
	"encoding/json"
	"go-rest-ws/server"
	"net/http"
)

type HomeResponse struct {
	Message string `json:"message"` // Serialized as "message"
	Status  bool   `json:"status"`
}

func HommeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		/*
			This is a sample response
			NewEncoder: Creates a new encoder that writes to w.
			Encode: Encodes the value and writes it to the stream.

			The response is a JSON object with two fields:
				message: "Hello World"
				status: "OK"

			The response is serialized as:
				{"message": "Hello World", "status": "OK"}
		*/
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Hello World",
			Status:  true,
		})
	}
}
