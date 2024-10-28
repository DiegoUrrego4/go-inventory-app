package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/DiegoUrrego4/inventory-app/config"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (s *Server) Run() {
	s.connectDatabase()

	r := s.registerRouter()

	log.Println("Starting server at port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func (s *Server) connectDatabase() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load project config %v", err)
	}

	db, err := sql.Open("postgres", "postgres://"+cfg.DBUser+":"+cfg.DBPassword+"@"+cfg.DBHost+":"+cfg.DBPort+"/"+cfg.DBName)
	if err != nil {
		log.Fatalf("Could not connect to database %v", err)
	}

	defer db.Close()

	log.Println("Connected to database")
}

func (s *Server) registerRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Inventory API!"))
	})

	return router
}
