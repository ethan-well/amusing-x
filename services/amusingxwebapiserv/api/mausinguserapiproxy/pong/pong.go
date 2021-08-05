package pong

import (
	"net/http"
)

func Pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("amusing user service"))
	return
}
