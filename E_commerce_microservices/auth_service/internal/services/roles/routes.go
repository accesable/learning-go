package roles

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"trann/ecom/auth_service/internal/types/payloads"
)

type RoleHandler struct {
	Service *RoleService
}

func NewRoleHandler(service *RoleService) *RoleHandler {
	return &RoleHandler{Service: service}
}

// Example: Creating a new role
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var input payloads.CreateRoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := h.Service.CreateRole(input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": role})
}

// Example: Assign a role to a user
func (h *RoleHandler) AssignRoleToUser(c *gin.Context) {
	var input payloads.AssignRoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.AssignRoleToUser(input.UserID, input.RoleName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role assigned successfully"})
}
