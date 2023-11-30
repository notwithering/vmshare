package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	filePath := flag.String("file", ".", "File or directory to be shared")
	port := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedPath := filepath.Join(*filePath, r.URL.Path)

		fileInfo, err := os.Stat(requestedPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if fileInfo.IsDir() {
			http.FileServer(http.Dir(requestedPath)).ServeHTTP(w, r)
		} else {
			http.ServeFile(w, r, requestedPath)
		}
	})

	address := fmt.Sprintf(":%d", *port)
	fmt.Printf("Sharing at http://127.0.0.1%s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
