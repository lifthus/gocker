package main

import (
	"flag"
	"gocker/internal/muxer"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "9595", "port number to run the server on(default: 9595)")
	flag.Parse()

	mux := http.NewServeMux()

	mux = muxer.SetGockerApi(mux)

	server := &http.Server{
		Addr:    ":" + *port,
		Handler: mux,
	}
	log.Println("Gocker server listening on port:", *port, "...")
	log.Fatal(server.ListenAndServe())
}
