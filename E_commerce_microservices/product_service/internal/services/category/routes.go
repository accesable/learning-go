package category

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"

	"trann/ecom/product_services/internal/types"
	"trann/ecom/product_services/internal/utils"
)

type Handler struct {
	store types.CategoryStore
}

func NewHandler(store types.CategoryStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoues(router *http.ServeMux) {
	router.HandleFunc("GET /api/v1/category", h.handleGetCategories)
	router.HandleFunc("POST /api/v1/category", h.handlePostCategory)
	router.HandleFunc("DELETE /api/v1/category/{id}", h.handleDeleteCategory)
	router.HandleFunc("GET /api/v1/category/{id}", h.handleGetCategoryById)
}

func (h *Handler) handleGetCategories(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetCategories(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) handlePostCategory(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateCategoryPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	// validate json payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	id, err := h.store.CreateCategory(r.Context(), &payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(
		w,
		http.StatusOK,
		map[string]string{"message": fmt.Sprintf("New Categoy Created with Id : %d", id)},
	)
}

func (h *Handler) handleDeleteCategory(w http.ResponseWriter, r *http.Request) {
	deletedId := r.PathValue("id")
	deletedIdInt, err := strconv.Atoi(deletedId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if _, err := h.store.GetCategoryById(r.Context(), deletedIdInt); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if err := h.store.DeleteCategory(r.Context(), deletedIdInt); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(
		w,
		http.StatusOK,
		map[string]string{
			"message": fmt.Sprintf("Category ID : %d Deleted Succesfully", deletedIdInt),
		},
	)
}

func (h *Handler) handleGetCategoryById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	selectedCategory, err := h.store.GetCategoryById(r.Context(), idInt)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, selectedCategory)
}
