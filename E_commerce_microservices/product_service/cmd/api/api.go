package api

import (
	"database/sql"
	"log"
	"net/http"

	mysqlc "trann/ecom/product_services/internal/model"
	"trann/ecom/product_services/internal/services/category"
	"trann/ecom/product_services/internal/services/items"
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
	queriesInstance := mysqlc.New(s.db)
	// Category handlers
	categoryStore := category.NewStore(queriesInstance)
	categoryHandlers := category.NewHandler(categoryStore)
	categoryHandlers.RegisterRoues(router)

	// Items handlers
	itemStore := items.NewStore(queriesInstance)
	itemHandlers := items.NewHandler(itemStore)
	itemHandlers.RegisterRoutes(router)
	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server started at %s", s.addr)

	return server.ListenAndServe()
}
