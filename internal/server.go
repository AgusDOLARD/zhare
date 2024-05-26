package internal

import (
	"net/http"
)

func Serve(addr string, fp string) error {
	srv := http.NewServeMux()

	fs := http.FileServer(http.Dir(fp))
	srv.Handle("/", fs)

	return http.ListenAndServe(addr, srv)
}
