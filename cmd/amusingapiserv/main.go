package main

import (
	"amusingx.fit/amusingx/services/amusingapiserv/api/pong"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/pong", pong.Pong)

	http.ListenAndServe(":8080", nil)
}