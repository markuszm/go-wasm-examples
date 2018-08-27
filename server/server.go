package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8080", "port")
	dir := flag.String("dir", "./out", "directory")
	flag.Parse()
	log.Fatal(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*dir))))
}
