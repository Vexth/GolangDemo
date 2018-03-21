package router

import (
	"io"
	"net/http"
)

// SayBye xxxxxxxx .
func SayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "byebye")
}
