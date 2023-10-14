package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"io"
	"net/http"
	"os"
)

const krampusVersion string = "1.2"

var fileUploadPath string = "./uploads" // Default upload path
var fileDownloadPath string = "./"      // Default download path
//var sslCertPath string = "./cert.pem"
//var sslKeyPath string = "./key.pem"

// File Uploads
func FileUpload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	os.MkdirAll(fileUploadPath, os.ModePerm)

	dst, err := os.Create(fmt.Sprintf("%s/%s", fileUploadPath, fileHeader.Filename))
	defer dst.Close()

	// Copy file[file] to destination[dst]
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("[POST] %s \n", fileHeader.Filename)
}

/*
// File Downloads
func FileDownload(w http.ResponseWriter, r *http.Request) {
  http.FileServer(http.Dir("."))
  //fmt.Printf("krampus(v1.1) starting at port: %d (SSL: %t)", variable)
}
*/
// HTTP/HTTPS Routing
func Routing(portChoice int16, sslChoice bool) {
	var portChoiceFormatted string = fmt.Sprintf(":%d", portChoice)
	FileDownload := http.FileServer(http.Dir(fileDownloadPath))

	http.HandleFunc("/upload", FileUpload) // http://0.0.0.0:<port>/upload
	http.Handle("/", FileDownload)         // http://0.0.0.0:<port>/

	if sslChoice == true {
		// Replace with sslCertPath / sslKeyPath functionality.
		http.ListenAndServeTLS(portChoiceFormatted, "./cert.pem", "./key.pem", nil)
	} else {
		http.ListenAndServe(portChoiceFormatted, nil)
	}
}

func main() {
	var portChoice int16 = 9001
	pflag.Int16Var(&portChoice, "port", portChoice, "port selection")
	var sslChoice bool = false
	pflag.BoolVar(&sslChoice, "ssl", sslChoice, "enables SSL, requires (exactly named) the 'cert.pem' and 'key.pem'")
	pflag.StringVar(&fileUploadPath, "file-upload-path", fileUploadPath, "file upload path")
	pflag.StringVar(&fileDownloadPath, "file-download-path", fileDownloadPath, "file serve path")

	pflag.Parse()

	fmt.Printf("krampus(v%s) starting at port: %d (SSL: %t)\n", krampusVersion, portChoice, sslChoice)
	Routing(portChoice, sslChoice)
}
