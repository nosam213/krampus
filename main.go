package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/quic-go/quic-go/http3"
	"github.com/spf13/pflag"
)

const krampusVersion string = "1.5"

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
var sslChoice bool
var quicChoice bool
var portChoice string = "9001"
var ipChoice string = "0.0.0.0"
var filesServe bool

// file upload handling
func FileUpload(w http.ResponseWriter, r *http.Request) {
	// GET
	if r.Method == "GET" {
		fmt.Fprint(w, htmlWebpage)
		return
	}

	// POST
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	os.MkdirAll(fileUploadPath, os.ModePerm) // makes directory if specified is not found

	dst, err := os.Create(fmt.Sprintf("%s/%s", fileUploadPath, fileHeader.Filename)) // Formats the location + filename
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	fmt.Printf("[POST] %s \n", fileHeader.Filename)
}

func main() {
	pflag.BoolVarP(&sslChoice, "tls", "e", false, "TLS (default @ ./cert.pem ./key.pem)")                      // -e , --tls
	pflag.BoolVarP(&quicChoice, "http3", "3", false, "enables QUIC/HTTP3 (UDP) (experimental)")                // -3 , --http3
	pflag.BoolVarP(&filesServe, "files", "f", false, "enable serving files")                                   // -f , --files
	pflag.StringVarP(&portChoice, "port", "p", portChoice, "port selection")                                   // -p, --port
	pflag.StringVarP(&ipChoice, "ip", "a", ipChoice, "ip selection")                                           // -a , --ip
	pflag.StringVar(&fileUploadPath, "file-upload-path", fileUploadPath, "file upload serve destination")      // --file-upload-path
	pflag.StringVar(&fileDownloadPath, "file-download-path", fileDownloadPath, "file download serve location") // --file-download-path
	pflag.StringVar(&sslCertPath, "ssl-cert-path", sslCertPath, "TLS certificate path")                        // --ssl-cert-path
	pflag.StringVar(&sslKeyPath, "ssl-key-path", sslKeyPath, "TLS key path")                                   // --ssl-key-path
	pflag.Parse()

	// routing
	var portChoiceFormatted string = fmt.Sprintf("%s:%s", ipChoice, portChoice)
	http.HandleFunc("/upload", FileUpload) // http://<address>:<port>/upload

	fmt.Printf("krampus(v%s) running at: %s [TLS: %t] [Files: %t]\n", krampusVersion, portChoiceFormatted, sslChoice, filesServe)

	// serving files disable by default
	if filesServe {
		FileDownload := http.FileServer(http.Dir(fileDownloadPath))
		http.Handle("/", FileDownload) // http://0.0.0.0:<port>/
	}

	if sslChoice {
		if quicChoice {
			err := http3.ListenAndServeQUIC(portChoiceFormatted, sslCertPath, sslKeyPath, nil)
			if err != nil {
				fmt.Printf("Couldn't bind to %s\n", portChoiceFormatted)
				os.Exit(1)
			}
		} else {
			err := http.ListenAndServeTLS(portChoiceFormatted, sslCertPath, sslKeyPath, nil)
			if err != nil {
				fmt.Printf("Couldn't bind to %s\n", portChoiceFormatted)
				os.Exit(1)
			}
		}
	} else {
		err := http.ListenAndServe(portChoiceFormatted, nil)
		if err != nil {
			fmt.Printf("Couldn't bind to %s\n", portChoiceFormatted)
			os.Exit(1)
		}
	}
}
