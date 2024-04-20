package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "")

	flag.Usage = func() {
		fmt.Println("vmshare [options...] directory")
		fmt.Println("  -port Port to listen on")
	}

	flag.Parse()

	dir := flag.Arg(0)
	if dir == "" {
		dir = "."
	}

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Sharing at http://127.0.0.1%s\n", addr)

	if err := http.ListenAndServe(addr, http.FileServer(http.Dir(dir))); err != nil {
		fmt.Println(err)
		return
	}
}
