package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"trann/ecom/auth_service/internal/services/roles"
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
	roleRepository := &roles.RoleRepositoryImpl{DB: s.db}
	roleService := &roles.RoleService{Repo: roleRepository}
	userRepository := &users.UserRepositoryImpl{DB: s.db}
	userService := &users.UserService{Repo: userRepository, RoleRepo: roleRepository}

	// Handler setup
	userHandler := users.NewUserHandler(userService, roleService)
	roleHandlers := roles.NewRoleHandler(roleService)

	r.POST("/signup", userHandler.Signup)
	r.POST("/signin", userHandler.Signin)
	r.POST("/assign-role", roleHandlers.AssignRoleToUser)
	r.POST("/create-role", roleHandlers.CreateRole)
	log.Printf("Server start at %s\n", s.addr)
	return r.Run(s.addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
