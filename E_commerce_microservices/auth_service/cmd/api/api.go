package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"trann/ecom/auth_service/internal/services/users"
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
	// Repository and Service setup
	userRepository := &users.UserRepositoryImpl{DB: s.db}
	userService := &users.UserService{Repo: userRepository}

	// Handler setup
	userHandler := users.NewUserHandler(userService)

	r.POST("/signup", userHandler.Signup)
	r.POST("/signin", userHandler.Signin)
	log.Printf("Server start at %s\n", s.addr)
	return r.Run(s.addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
