package payloads

type SignupInput struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type SigninInput struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
type UserResponse struct {
	ID       uint     `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
}

// CreateRoleInput represents the input for creating a new role
type CreateRoleInput struct {
	Name string `json:"name" binding:"required,min=3,max=50"`
}

// AssignRoleInput represents the input for assigning a role to a user
type AssignRoleInput struct {
	UserID   uint   `json:"userId"   binding:"required"`
	RoleName string `json:"roleName" binding:"required,min=3,max=50"`
}
