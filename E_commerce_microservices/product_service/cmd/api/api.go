package api

import (
	"database/sql"
	"log"
	"net/http"

	"trann/ecom/product_services/internal/services/category"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	categoryStore := category.NewStore(s.db)
	categoryHandlers := category.NewHandler(categoryStore)
	categoryHandlers.RegisterRoues(router)
	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server started at %s", s.addr)

	return server.ListenAndServe()
}
