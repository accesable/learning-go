package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"trann/ecom/auth_service/internal/types/payloads"
)

// UserHandler handles HTTP requests related to user actions
type UserHandler struct {
	Service *UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// Validate request payloads
var validate = validator.New()

// Signup handles user registration
func (h *UserHandler) Signup(c *gin.Context) {
	var input payloads.SignupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.Signup(input.Username, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT token (for example purposes, you'll need to implement this)
	token, err := h.Service.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	response := payloads.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	c.JSON(http.StatusOK, gin.H{"authToken": token, "user": response})
}

// Signin handles user login
func (h *UserHandler) Signin(c *gin.Context) {
	var input payloads.SigninInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.Signin(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token (for example purposes, you'll need to implement this)
	token, err := h.Service.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	response := payloads.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": response})
}
