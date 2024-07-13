package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"trann/ecom/order_service/internal/services/order"
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
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	hanlders := order.NewHandler()
	hanlders.RegisterRoutes(r)
	log.Printf("Server start at %s\n", s.addr)
	return r.Run(s.addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
