package main

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"os"

	go_qr "github.com/piglig/go-qr"
)

const (
	port       = 3000
	blackBlock = "\033[40m  \033[0m"
	whiteBlock = "\033[47m  \033[0m"
)

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}
	fp := os.Args[1]

	localIP, err := getLocalIP()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	url := fmt.Sprintf("http://%s:%d", localIP, port)
	err = generateQR(url)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Printf("Serving on: %s", url)

	err = serveFile(fp)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func serveFile(fp string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open(fp)
		if err != nil {
			http.Error(w, "Error opening file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().
			Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileInfo.Name()))
		http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)

	})
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func generateQR(content string) error {
	qr, err := go_qr.EncodeText(content, go_qr.Low)
	if err != nil {
		return err
	}
	buf := bytes.Buffer{}
	border := 4
	for y := -border; y < qr.GetSize()+border; y++ {
		for x := -border; x < qr.GetSize()+border; x++ {
			if !qr.GetModule(x, y) {
				buf.WriteString(blackBlock)
			} else {
				buf.WriteString(whiteBlock)
			}
		}
		buf.WriteString("\n")
	}
	fmt.Print(buf.String())
	return nil
}

func usage() {
	fmt.Printf("Usage: %s FILEPATH\n", os.Args[0])
}

func getLocalIP() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		// Skip loopback and down interfaces
		if iface.Flags&net.FlagLoopback == 0 && iface.Flags&net.FlagUp != 0 {
			addrs, err := iface.Addrs()
			if err != nil {
				return "", err
			}

			// Prefer IPv4 addresses
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String(), nil
					}
				}
			}
		}
	}

	return "", fmt.Errorf("No suitable local IP address found")
}
