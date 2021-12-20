package comm

import (
	"context"
	"net/http"
)

func Ctx(r *http.Request) context.Context {
	return context.WithValue(context.Background(), "reqID", "1111")
}
