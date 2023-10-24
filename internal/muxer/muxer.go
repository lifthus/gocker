package muxer

import (
	"net/http"
	"os"
)

func SetGockerApi(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the Gocker server!" + os.Getenv("ENV1")))
	})
	return mux
}
