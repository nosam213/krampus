package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"io"
	"net/http"
	"os"
)

/*
[X] Implement port selection through flags.
[O] Pass the 'dirChoice' variable to FileUpload through handler.
[X] Implement SSL trhough http.ListenAndServeTLS function. & if/else
[O] Implement basic password authentication.
[O] Implement a shit ton more logging for fatal errors, http module related
*/

// File Uploading
func FileUpload(w http.ResponseWriter, r *http.Request) {
	// FormFile to match arguments
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	// Creates uploads directory if not present
	os.MkdirAll("./uploads", os.ModePerm)

	// Create empty file 
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
func Routing(portChoice int16, sslChoice bool) {
	// Formatting 'portChoice' to look like ex: ":8080" for http
	var portChoiceFormatted string = fmt.Sprintf(":%d", portChoice)
	
	http.HandleFunc("/upload", FileUpload)
	
	if sslChoice == true {
		http.ListenAndServeTLS(portChoiceFormatted, "cert.pem", "key.pem", nil)
	} else {
		http.ListenAndServe(portChoiceFormatted, nil)
	}
}

func main() {
	// Flag Parsing
	var portChoice int16 = 9001
	pflag.Int16Var(&portChoice, "port", portChoice, "port selection (default 9001)")
	/*
	var dirChoice string = "./uploads"
	pflag.StringVar(&dirChoice, "dir", dirChoice, "directory selection (default ./uploads")
	*/
	var sslChoice bool = false
	pflag.BoolVar(&sslChoice, "ssl", sslChoice, "enables SSL, requires (exactly named) the 'cert.pem' and 'key.pem'")
	
	pflag.Parse()
	
	fmt.Printf("Server starting at port: %d", portChoice)
	Routing(portChoice, sslChoice)
}