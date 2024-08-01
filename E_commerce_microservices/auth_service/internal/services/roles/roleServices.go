package roles

import (
	"trann/ecom/auth_service/internal/models"
	"trann/ecom/auth_service/internal/types/repositories"
)

type RoleService struct {
	Repo repositories.RoleRepository
}

func (s *RoleService) CreateRole(name string) (*models.Role, error) {
	role := &models.Role{RoleName: name}
	if err := s.Repo.CreateRole(role); err != nil {
		return nil, err
	}
	return role, nil
}

func (s *RoleService) GetRoleByID(id uint) (*models.Role, error) {
	return s.Repo.GetRoleByID(id)
}

func (s *RoleService) GetRoleByName(name string) (*models.Role, error) {
	return s.Repo.GetRoleByName(name)
}

func (s *RoleService) GetAllRoles() ([]models.Role, error) {
	return s.Repo.GetAllRoles()
}

func (s *RoleService) AssignRoleToUser(userID uint, roleName string) error {
	role, err := s.Repo.GetRoleByName(roleName)
	if err != nil {
		return err
	}
	return s.Repo.AssignRoleToUser(userID, role.ID)
}

func (s *RoleService) GetUserRoles(userID uint) ([]models.Role, error) {
	return s.Repo.GetUserRoles(userID)
}
