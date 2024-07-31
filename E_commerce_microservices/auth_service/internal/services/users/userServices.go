package users

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"trann/ecom/auth_service/internal/models"
	"trann/ecom/auth_service/internal/types/repositories"
)

type UserService struct {
	Repo repositories.UserRepository
}

func (s *UserService) Signup(username, email, password string) (*models.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(passwordHash),
	}

	err = s.Repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Signin(email, password string) (*models.User, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GenerateJWT(user *models.User) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("JWT secret is not set")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	log.Printf("jwt %v", jwtSecret)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
