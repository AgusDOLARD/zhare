package server

import (
	"io"
	"log/slog"
	"net/http"
)

type Server struct {
	addr   string
	reader io.Reader
}

func NewServer(addr string, data io.Reader) *Server {
	return &Server{
		addr:   addr,
		reader: data,
	}
}

func (s *Server) Start() error {
	slog.Info("Starting server", "addr", s.addr)
	http.HandleFunc("/", s.handleFileDownload)
	return http.ListenAndServe(s.addr, nil)
}

func (s *Server) handleFileDownload(w http.ResponseWriter, r *http.Request) {
	b, err := io.Copy(w, s.reader)
	if err != nil {
		slog.Error("Error serving file", "err", err)
		return
	}
	slog.Info("Served file", "path", r.URL.Path, "size", b)
}
