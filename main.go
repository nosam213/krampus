package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// File Uploading
func FileUpload(w http.ResponseWriter, r *http.Request) {
	// FormFile to match arguments
	file, fileHeader, err := r.FormFile("file")

	defer file.Close()

	// Creates uploads directory if not present
	os.MkdirAll("./uploads", os.ModePerm)

	// Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("./uploads/%s", fileHeader.Filename))

	defer dst.Close()

	// Copy file to destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// HTTP Routing
func Routing() {
	http.HandleFunc("/upload", FileUpload)
	http.ListenAndServe(":9001", nil)
}

func main() {
	fmt.Println("Server starting at port: 9001")
	Routing()
}
