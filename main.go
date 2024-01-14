package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"io"
	"net/http"
	"os"
)

const krampusVersion string = "1.3"

const htmlWebpage string = `
<!DOCTYPE html>
<html>
<head>
  <title>Krampus File Upload</title>
</head>
<body>
  <form method="post" enctype="multipart/form-data">
  <label for="file">File</label>
    <input id="file" name="file" type="file"></input>
  <button>Upload</button>
  </form>
</body>
</html>`

var fileUploadPath string = "./uploads"
var fileDownloadPath string = "./"
var sslCertPath string = "./cert.pem"
var sslKeyPath string = "./key.pem"

// File Uploads
func FileUpload(w http.ResponseWriter, r *http.Request) {
	//-- [GET] --//
	if r.Method == "GET" {
		fmt.Fprintf(w, htmlWebpage)
		return
	}
  
  //-- [POST] --//
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	os.MkdirAll(fileUploadPath, os.ModePerm) // Makes directory if specified not found

	dst, err := os.Create(fmt.Sprintf("%s/%s", fileUploadPath, fileHeader.Filename)) // Formats the location + filename
	defer dst.Close()

	// Copy file[file] to destination[dst]
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("[POST] %s \n", fileHeader.Filename)
}

// Routing
func Routing(portChoice int16, sslChoice bool) {
	var portChoiceFormatted string = fmt.Sprintf(":%d", portChoice)
	FileDownload := http.FileServer(http.Dir(fileDownloadPath))

	http.HandleFunc("/upload", FileUpload) // http://0.0.0.0:<port>/upload
	http.Handle("/", FileDownload)         // http://0.0.0.0:<port>/

	if sslChoice == true {
		// Replace with sslCertPath / sslKeyPath functionality.
		http.ListenAndServeTLS(portChoiceFormatted, sslCertPath, sslKeyPath, nil)
	} else {
		http.ListenAndServe(portChoiceFormatted, nil)
	}
}

func main() {
	var portChoice int16 = 9001
	pflag.Int16Var(&portChoice, "port", portChoice, "port selection")
	var sslChoice bool = false
	pflag.BoolVar(&sslChoice, "ssl", sslChoice, "enables HTTPS (default ./cert.pem and ./key.pem)") // --ssl
	pflag.StringVar(&fileUploadPath, "file-upload-path", fileUploadPath, "file upload path")        // --file-upload-path
	pflag.StringVar(&fileDownloadPath, "file-download-path", fileDownloadPath, "file serve path")   // --file-download-path
	pflag.StringVar(&sslCertPath, "ssl-cert-path", sslCertPath, "TLS certificate path")             // --ssl-cert-path
	pflag.StringVar(&sslKeyPath, "ssl-key-path", sslKeyPath, "TLS key path")                        // --ssl-key-path

	pflag.Parse()

	fmt.Printf("krampus(v%s) starting at port: %d (SSL: %t)\n", krampusVersion, portChoice, sslChoice)
	Routing(portChoice, sslChoice)
}
