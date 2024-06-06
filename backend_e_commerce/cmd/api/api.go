package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/trann/e_commerce/internal/service/product"
	"github.com/trann/e_commerce/internal/service/user"
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

	// userHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	userID := r.PathValue("userID")
	// 	w.Write([]byte("User ID: " + userID))
	// })
	// router.Handle("GET /api/v1/user/{userID}", userHandler)

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(router)

	productStore := product.NewStore(s.db)
	productHandlers := product.NewHandler(productStore, userStore)
	productHandlers.RegisterRoutes(router)
	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server started at %s", s.addr)

	return server.ListenAndServe()
}
