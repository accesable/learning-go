package category

import (
	"net/http"
	"trann/ecom/product_services/internal/types"
	"trann/ecom/product_services/internal/utils"
)


type Handler struct{
  store types.CategoryStore
}

func NewHandler(store types.CategoryStore)  *Handler{
  return &Handler{
    store: store,
  } 
}

func (h *Handler)RegisterRoues(router *http.ServeMux){
  router.HandleFunc("GET /api/v1/category",h.handleGetCategories)
}

func (h *Handler)handleGetCategories(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetCategories(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, products)
}
