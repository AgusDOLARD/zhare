package server

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/*
var templates embed.FS

const uploadHTMLPage = "templates/upload.html"

func Serve(addr string, fp ...string) error {
	srv := http.NewServeMux()

	srv.HandleFunc("/upload", postUploadHandle)

	srv.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if len(fp) == 1 {
			getUploadHandle(w, r)
		} else {
			getDownloadHandle(w, r, fp[0])
		}
	})

	return http.ListenAndServe(addr, srv)
}

func postUploadHandle(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	dst, err := os.Create(filepath.Join("./", header.Filename))
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// upload the file to destination path
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getUploadHandle(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(templates, uploadHTMLPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getDownloadHandle(w http.ResponseWriter, r *http.Request, fp string) {
	file, err := os.Open(fp)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().
		Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileInfo.Name()))
	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)

}
