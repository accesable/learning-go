package user

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/trann/e_commerce/internal/types"
)

func TestUserServiceHanlders(t *testing.T) {
	userStore := &mockUserStore{}
	hanlder := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Nhut Anh",
			LastName:  "Tran",
			Email:     "example",
			Password:  "asd",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("POST /api/v1/register", hanlder.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d got, %d",
				http.StatusBadRequest, rr.Code)
		}
	})
	t.Run("should fail if the user payload is valid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Nhut Anh",
			LastName:  "Tran",
			Email:     "example@gmail.com",
			Password:  "asd",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("POST /api/v1/register", hanlder.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d got, %d",
				http.StatusCreated, rr.Code)
		}
	})
}

type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	return nil, fmt.Errorf("User not founded")
}

func (m *mockUserStore) GetUserByID(ctx context.Context, id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(ctx context.Context, user types.User) error {
	return nil
}
