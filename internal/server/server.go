package server

import (
	"net/http"
)

func Serve(addr string, fp ...string) error {
	srv := http.NewServeMux()
	srv.HandleFunc("/upload", PostUploadHandle)
	srv.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if len(fp) == 1 {
			GetUploadHandle(w, r)
		} else {
			GetDownloadHandle(w, r, fp[0])
		}
	})
	return http.ListenAndServe(addr, srv)
}
