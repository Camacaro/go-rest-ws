package main

import (
	"context"
	"go-rest-ws/handlers"
	"go-rest-ws/middlewares"
	"go-rest-ws/server"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar las variables de entorno
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	// Create a new server
	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

// func BindRoutes(s server.Server, r *mux.Router) {

// 	api := r.PathPrefix("/api/v1").Subrouter()

// 	// Create a new Hub, websocket
// 	// hub := websocket.NewHub() -> Se crea un nuevo hub en el server.go

// 	// Middlewares
// 	// Para cada ruta que esta aqui pasara por este middleware
// 	r.Use(middlewares.CheckAuthMiddleware(s))

// 	r.HandleFunc("/", handlers.HommeHandler(s)).Methods("GET")
// 	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods("POST")
// 	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
// 	r.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)
// 	r.HandleFunc("/posts", handlers.InsertPostHandler(s)).Methods(http.MethodPost)
// 	r.HandleFunc("/posts/{id}", handlers.GetPostByIdHandler(s)).Methods(http.MethodGet)
// 	r.HandleFunc("/posts/{id}", handlers.UpdatePostHandler(s)).Methods(http.MethodPut)
// 	r.HandleFunc("/posts/{id}", handlers.DeletePostHandler(s)).Methods(http.MethodDelete)
// 	r.HandleFunc("/posts", handlers.ListPostHandler(s)).Methods(http.MethodGet)

// 	// Ruta qeu manejara el websocket, y activara el hub
// 	// go hub.Run() -> Ya no se usa aqui sino en server.go
// 	// r.HandleFunc("/ws", hub.HandleWebScoket)
// 	r.HandleFunc("/ws", s.Hub().HandleWebScoket)
// }

func BindRoutes(s server.Server, r *mux.Router) {

	// Sub rutas, para que esten protegidas por token
	// Las que no lo tengan se puede acceder sin el token
	api := r.PathPrefix("/api/v1").Subrouter()

	// Create a new Hub, websocket
	// hub := websocket.NewHub() -> Se crea un nuevo hub en el server.go

	// Middlewares
	// Para cada ruta que esta aqui pasara por este middleware
	api.Use(middlewares.CheckAuthMiddleware(s))

	r.HandleFunc("/", handlers.HommeHandler(s)).Methods("GET")
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)
	api.HandleFunc("/posts", handlers.InsertPostHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/posts/{id}", handlers.GetPostByIdHandler(s)).Methods(http.MethodGet)
	api.HandleFunc("/posts/{id}", handlers.UpdatePostHandler(s)).Methods(http.MethodPut)
	api.HandleFunc("/posts/{id}", handlers.DeletePostHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/posts", handlers.ListPostHandler(s)).Methods(http.MethodGet)

	// Ruta qeu manejara el websocket, y activara el hub
	// go hub.Run() -> Ya no se usa aqui sino en server.go
	// r.HandleFunc("/ws", hub.HandleWebScoket)
	r.HandleFunc("/ws", s.Hub().HandleWebScoket)
}
