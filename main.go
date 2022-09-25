package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"io"
	"net/http"
	"os"
)

// File Uploading
func FileUpload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	os.MkdirAll("./uploads", os.ModePerm)

	dst, err := os.Create(fmt.Sprintf("./uploads/%s", fileHeader.Filename))
	defer dst.Close()

	// Copy file[file] to destination[dst]
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// HTTP/HTTPS Routing
func Routing(portChoice int16, sslChoice bool) {
	var portChoiceFormatted string = fmt.Sprintf(":%d", portChoice)
	FileDownload := http.FileServer(http.Dir("."))
	
	http.HandleFunc("/upload", FileUpload)
	http.Handle("/", FileDownload)
	
	if sslChoice == true {
		http.ListenAndServeTLS(portChoiceFormatted, "cert.pem", "key.pem", nil)
	} else {
		http.ListenAndServe(portChoiceFormatted, nil)
	}
}

func main() {
	var portChoice int16 = 9001
	pflag.Int16Var(&portChoice, "port", portChoice, "port selection (default 9001)")
	var sslChoice bool = false
	pflag.BoolVar(&sslChoice, "ssl", sslChoice, "enables SSL, requires (exactly named) the 'cert.pem' and 'key.pem'")
	
	pflag.Parse()
	
	fmt.Printf("krampus(v1.1) starting at port: %d (SSL: %t)", portChoice, sslChoice)
	Routing(portChoice, sslChoice)
}
