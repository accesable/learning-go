package repositories

import "trann/ecom/auth_service/internal/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}

// RoleRepository defines the methods for working with roles
type RoleRepository interface {
	CreateRole(role *models.Role) error
	GetRoleByID(id uint) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]models.Role, error)
	AssignRoleToUser(userID uint, roleID uint) error
	GetUserRoles(userID uint) ([]models.Role, error)
}
