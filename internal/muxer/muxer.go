package muxer

import "net/http"

func SetGockerApi(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the Gocker server"))
	})
	return mux
}
