package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	PORT string
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := os.CreateTemp("uploads", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes(port string) {
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})
	http.ListenAndServe(":"+port, nil)
}

func main() {
	if PORT = os.Getenv("APP_PORT"); PORT == "" {
		PORT = "8081"
	}
	fmt.Println("App started .. up and running on PORT:", PORT)
	setupRoutes(PORT)
}
