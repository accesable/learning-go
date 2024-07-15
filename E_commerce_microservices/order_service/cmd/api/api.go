package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"trann/ecom/order_service/internal/services/order"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	r := gin.Default()

	orderRepository := order.NewStore(s.db)
	hanlders := order.NewHandler(orderRepository)
	hanlders.RegisterRoutes(r)
	log.Printf("Server start at %s\n", s.addr)
	return r.Run(s.addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
