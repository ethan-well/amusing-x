package home

import "net/http"

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello index"))
	return
}
