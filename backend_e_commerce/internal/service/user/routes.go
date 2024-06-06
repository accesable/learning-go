package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/trann/e_commerce/internal/config"
	"github.com/trann/e_commerce/internal/service/auth"
	"github.com/trann/e_commerce/internal/types"
	"github.com/trann/e_commerce/internal/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /api/v1/login", h.handleLogin)
	router.HandleFunc("POST /api/v1/register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Get Json Payload
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// validate json payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}
	// check if user exist
	user, err := h.store.GetUserByEmail(r.Context(), payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already is not existed", payload.Email))
		return
	}
	if err := auth.ComparePassword(user.Password, payload.Password); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid password with email %s", payload.Email))
		return
	}
	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Get Json Payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// validate json payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}
	// check if user exist
	_, err := h.store.GetUserByEmail(r.Context(), payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already existed", payload.Email))
		return
	}

	hashedPassword, _ := auth.HashedPassword(payload.Password)
	// if it does not we create the new user
	err = h.store.CreateUser(r.Context(), types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, nil)
}
