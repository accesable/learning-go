package server

import (
	"log"
	"net/http"

	"trann/ecom/file_server_storage/internal/utils/handlers"
)

type FileServer struct {
	addr      string
	uploadDir string
}

func NewFileServer(addr string, uploadDir string) *FileServer {
	return &FileServer{
		addr:      addr,
		uploadDir: uploadDir,
	}
}

func (s *FileServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("POST /upload/{dirName}/{id}", handlers.UploadHandler)

	fs := http.FileServer(http.Dir("uploads"))

	router.Handle("/", fs)
	srv := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server started at %s ", s.addr)

	return srv.ListenAndServe()
}
