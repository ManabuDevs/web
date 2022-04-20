package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed app/build
var embeddedFiles embed.FS

func main() {
	fmt.Println("Starting Server")
	http.Handle("/", http.FileServer(getFileSystem()))
	http.ListenAndServe(":9000", nil)
}

func getFileSystem() http.FileSystem {

	// Get the build subdirectory as the
	// root directory so that it can be passed
	// to the http.FileServer
	fsys, err := fs.Sub(embeddedFiles, "app/build")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
