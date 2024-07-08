package main

import (
	"log"

	"trann/ecom/file_server_storage/cmd/server"
)

func main() {
	srv := server.NewFileServer(":8089", "./uploads")
	if err := srv.Run(); err != nil {
		log.Fatal("Error Running Server ", err)
	}
}
