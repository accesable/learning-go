package items

import (
	"net/http"

	"trann/ecom/product_services/internal/types"
	"trann/ecom/product_services/internal/utils"
)

type Handler struct {
	store types.ItemStore
}

func NewHandler(store types.ItemStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /api/v1/items", h.getItems)
}

func (h *Handler) getItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.GetItems(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, items)
}
