package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

//go:embed dist
var content embed.FS

func main() {
	http.Handle("/", http.FileServer(getFileSystem()))

	fmt.Printf("Starting server at port 8083\n")
	if err := http.ListenAndServe(":8083", nil); err != nil {
		log.Fatal(err)
	}
}

func getFileSystem() http.FileSystem {
	// Get the build subdirectory as the
	// root directory so that it can be passed
	// to the http.FileServer
	fsys, err := fs.Sub(content, "dist")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

