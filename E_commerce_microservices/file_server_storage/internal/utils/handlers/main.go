package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"trann/ecom/file_server_storage/internal/utils"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form in the request, setting a max memory limit
	r.ParseMultipartForm(10 << 20) // 10 MB limit

	// get dir name
	dirName := r.PathValue("dirName")
	id := r.PathValue("id")
	// fileName := r.URL.Query().Get("fileName")

	// Retrieve the file from the form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error retrieving file:", err)
		http.Error(w, "Error retrieving file", http.StatusInternalServerError)
		return
	}
	defer file.Close() // Ensure the file is closed after we're done

	// Print file information for debugging purposes
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// ext := filepath.Ext(handler.Filename)
	// Create a new file in the local filesystem
	// dst, err := os.Create("./uploads/uploaded_" + handler.Filename)
	dst, err := os.Create(fmt.Sprintf("./uploads/%s/%s/%s", dirName, id, handler.Filename))
	if err != nil {
		fmt.Println("Error creating file:", err)
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer dst.Close() // Ensure the destination file is closed after we're done

	// Copy the uploaded file's data to the destination file
	_, err = io.Copy(dst, file) // Stream the data from the reader to the writer
	if err != nil {
		fmt.Println("Error saving file:", err)
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	// Inform the client that the file has been uploaded successfully
	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"msg":  fmt.Sprintf("image uploaded at path /%s/%s", dirName, handler.Filename),
		"path": fmt.Sprintf("/%s/%s/%s", dirName, id, handler.Filename),
	})
}
