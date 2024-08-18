package server

import (
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

//go:embed web/index.html
var indexHTML string

type Server struct {
	addr  string
	files []File
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

func NewServer(addr string, files []string) *Server {
	f := make([]File, 0, len(files))
	for _, path := range files {
		f = append(f, *NewFile(path))
	}
	return &Server{
		addr:  addr,
		files: f,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("GET /", s.indexHanlder)
	http.HandleFunc("GET /file/{id}", s.fileHandler)
	return http.ListenAndServe(s.addr, nil)
}

func (s *Server) indexHanlder(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.New("index").
		Funcs(funcMap).
		Parse(indexHTML)
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
