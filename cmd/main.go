package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type IndexPage struct{}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading files\n")

	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("input-file")
	if err != nil {
		log.Fatalf("Error occured during file upload. Error: %s", err.Error())
		return
	}
	defer file.Close()
	fmt.Printf("File Name: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)

	tempFile, err := ioutil.TempFile("bin/temp-files", "*.xlsx")
	if err != nil {
		log.Fatalf("Error occured during file save. Error: %s", err.Error())
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Error occured during read file. Error: %s", err.Error())
	}
	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "Successfuly upload the file\n")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := IndexPage{}
	t, _ := template.ParseFiles("client/templates/index.html")
	t.Execute(w, p)
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Go files Upload!")
	setupRoutes()
}
