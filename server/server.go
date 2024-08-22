package server

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

//go:embed web/index.html web/upload.html
var pages embed.FS

type Server struct {
	addr   string
	files  []File
	logger *slog.Logger
}

type ServerOpts struct {
	Addr         string
	Files        []string
	EnableLogger bool
}

type File struct {
	ID   int
	Path string
}

func NewFile(path string) *File {
	return &File{
		ID:   rand.Int(),
		Path: path,
	}
}

func NewServer(opts *ServerOpts) *Server {
	var (
		f      = make([]File, 0, len(opts.Files))
		logger = slog.Default()
	)

	for _, path := range opts.Files {
		f = append(f, *NewFile(path))
	}

	if !opts.EnableLogger {
		logHandler := slog.NewTextHandler(io.Discard, &slog.HandlerOptions{})
		logger = slog.New(logHandler)
	}

	return &Server{
		addr:   opts.Addr,
		files:  f,
		logger: logger,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", s.indexHanlder)
	mux.HandleFunc("GET /upload", s.uploadHandler)
	mux.HandleFunc("POST /upload", s.uploadPostHandler)
	mux.HandleFunc("GET /file/{id}", s.fileHandler)

	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.LoggingMiddleware(mux),
	}
	s.logger.Info("Starting server", "addr", s.addr)
	return srv.ListenAndServe()
}

func (s *Server) indexHanlder(w http.ResponseWriter, r *http.Request) {
	fileLen := len(s.files)
	if fileLen == 0 {
		http.Redirect(w, r, "/upload", http.StatusSeeOther)
		return
	}
	indexHTML, err := pages.ReadFile("web/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("index").
		Funcs(funcMap).
		Parse(string(indexHTML))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, s.files)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) fileHandler(w http.ResponseWriter, r *http.Request) {
	fileID, _ := strconv.Atoi(r.PathValue("id"))
	file, err := s.getFile(fileID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	io.Copy(w, file)
}

func (s *Server) uploadHandler(w http.ResponseWriter, r *http.Request) {
	uploadHTML, err := pages.ReadFile("web/upload.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("upload").
		Funcs(funcMap).
		Parse(string(uploadHTML))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) uploadPostHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	destFile, err := os.Create(header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File %s uploaded successfully", header.Filename)
}

func (s *Server) getFile(id int) (*os.File, error) {
	for _, f := range s.files {
		if f.ID == id {
			return os.Open(f.Path)
		}
	}
	return nil, fmt.Errorf("file not found")
}

// template helper functions
var funcMap = template.FuncMap{
	"name": func(s string) string {
		return filepath.Base(s)
	},
}
