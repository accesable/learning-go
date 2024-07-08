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
	log.Println("Injected Queries to Category repository")
	categoryStore := category.NewStore(queriesInstance)
	categoryHandlers := category.NewHandler(categoryStore)
	categoryHandlers.RegisterRoues(router)

	// Items handlers
	log.Println("Injected Queries to Items repository")
	itemStore := items.NewStore(queriesInstance)
	itemHandlers := items.NewHandler(itemStore, categoryStore)
	itemHandlers.RegisterRoutes(router)
	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server started at %s", s.addr)

	return server.ListenAndServe()
}
