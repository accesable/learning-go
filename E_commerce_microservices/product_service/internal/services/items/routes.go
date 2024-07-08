package items

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"

	"trann/ecom/product_services/internal/types"
	"trann/ecom/product_services/internal/utils"
)

type Handler struct {
	store         types.ItemStore
	categoryStore types.CategoryStore
}

func NewHandler(store types.ItemStore, categoryStore types.CategoryStore) *Handler {
	return &Handler{
		store:         store,
		categoryStore: categoryStore,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /api/v1/items", h.handleGetItems)
	router.HandleFunc("POST /api/v1/items", h.handlePostItem)
	router.HandleFunc("DELETE /api/v1/items/{id}", h.handleDeleteItem)
}

func (h *Handler) handleGetItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.GetItems(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, items)
}

func (h *Handler) handlePostItem(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateItemPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload : %v", errors))
		return
	}
	// check if category id existed
	_, err := h.categoryStore.GetCategoryById(r.Context(), int(payload.CategoryID))
	if err != nil {
		utils.WriteError(
			w,
			http.StatusInternalServerError,
			fmt.Errorf("Error finding category Id : %v", payload.CategoryID),
		)
		return
	}
	id, err := h.store.CreateItem(r.Context(), &payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(
		w,
		http.StatusOK,
		map[string]string{"message": fmt.Sprintf("New Item Created Id : %v", id)},
	)
}

func (h *Handler) handleDeleteItem(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	id, err := strconv.Atoi(idPath)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	deletedRows, err := h.store.DeleteItem(r.Context(), int64(id))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if deletedRows == 0 {
		utils.WriteJSON(w, http.StatusNotFound, map[string]string{
			"message": fmt.Sprintf(
				"Item ID : %v is not deleted (maybe is not existed or already deleted)",
				idPath,
			),
		})
		return
	}
	utils.WriteJSON(
		w,
		http.StatusOK,
		map[string]string{
			"message": fmt.Sprintf("Item ID : %v Deleted Succesfully", idPath),
		},
	)
}
