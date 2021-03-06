package main

import (
	"filesStoreProject/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("file/upload/suc", handler.UploadSuncHandler)
	http.HandleFunc("file/meta", handler.GetFileMetaHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
}
