package roles

import (
	"gorm.io/gorm"

	"trann/ecom/auth_service/internal/models"
)

type RoleRepositoryImpl struct {
	DB *gorm.DB
}

func (r *RoleRepositoryImpl) CreateRole(role *models.Role) error {
	return r.DB.Create(role).Error
}

func (r *RoleRepositoryImpl) GetRoleByID(id uint) (*models.Role, error) {
	var role models.Role
	if err := r.DB.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepositoryImpl) GetRoleByName(name string) (*models.Role, error) {
	var role models.Role
	if err := r.DB.Where("role_name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepositoryImpl) GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	if err := r.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RoleRepositoryImpl) AssignRoleToUser(userID uint, roleID uint) error {
	user := models.User{ID: userID}
	role := models.Role{ID: roleID}
	return r.DB.Model(&user).Association("Roles").Append(&role)
}

func (r *RoleRepositoryImpl) GetUserRoles(userID uint) ([]models.Role, error) {
	var user models.User
	if err := r.DB.Preload("Roles").First(&user, userID).Error; err != nil {
		return nil, err
	}
	return user.Roles, nil
}
