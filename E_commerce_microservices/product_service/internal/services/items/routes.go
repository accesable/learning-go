package items

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"

	"trann/ecom/product_services/internal/config"
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
	router.HandleFunc("POST /multipart/v1/items/{id}", h.handlePostImageToItemOnId)
	router.HandleFunc("PATCH /api/v1/items/{id}", h.handlePatchItem)
}

func (h *Handler) handlePatchItem(w http.ResponseWriter, r *http.Request) {
	var updatePayload types.PartialUpdateItem
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if err := utils.ParseJSON(r, &updatePayload); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if err := utils.Validate.Struct(updatePayload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload : %v", errors))
		return
	}
	affectedRows, err := h.store.UpdateItemById(r.Context(), id, updatePayload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	log.Printf("%d rows changed", affectedRows)
	utils.WriteJSON(
		w,
		http.StatusOK,
		map[string]string{"msg": fmt.Sprintf("Updated Item Id : %v Succesfully", id)},
	)
}

func (h *Handler) handleGetItems(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	var options []types.GetItemsOption
	if showCategoryName := queries.Get("showCategoryName"); showCategoryName == "true" {
		options = append(options, types.WithCategoryName())
	}

	if includeImgURLs := queries.Get("includeImgURLs"); includeImgURLs == "true" {
		options = append(options, types.WithImgURLs())
	}
	items, err := h.store.GetItems(r.Context(), options...)
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

func (h *Handler) handlePostImageToItemOnId(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // 10 MB limit
	idStr := r.PathValue("id")
	itemId, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(
			w,
			http.StatusInternalServerError,
			fmt.Errorf("Error parsing id to number"),
		)
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		utils.WriteError(
			w,
			http.StatusInternalServerError,
			fmt.Errorf("Error retriving file"),
		)
		return
	}
	defer file.Close() // Ensure the file is closed after we're done

	// Print file information for debugging purposes
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	uploadResponse, err := uploadFileToServer(file, handler.Filename, idStr)
	if err != nil {
		utils.WriteError(
			w,
			http.StatusInternalServerError,
			fmt.Errorf("Error uploading file: %v", err),
		)
		return
	}
	id, err := h.store.UploadImageToItemId(r.Context(), types.ItemImage{
		ItemID:      int64(itemId),
		ImageUrl:    uploadResponse.Path,
		DisplayName: handler.Filename,
	})
	if err != nil {
		utils.WriteError(
			w,
			http.StatusInternalServerError,
			fmt.Errorf("error inserting new file record : %v", err),
		)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"msg": fmt.Sprintf("file upload Succesfully with file id : %d", id),
	})
}

// Struct to match the expected JSON response format
type UploadResponse struct {
	Msg  string `json:"msg"`
	Path string `json:"path"`
}

func uploadFileToServer(
	file multipart.File,
	fileName string,
	idStr string,
) (*UploadResponse, error) {
	// Create a new multipart writer
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Create the file field
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}

	// Copy the file data to the multipart writer
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	// Close the writer
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// Create a new POST request
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/upload/items/%s", config.Envs.FILE_SERVER_CONFIG, idStr),
		body,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"failed to upload file, server responded with status code %d",
			resp.StatusCode,
		)
	}

	// Read and decode the JSON response
	var uploadResponse UploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResponse); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return &uploadResponse, nil
}
