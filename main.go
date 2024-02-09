package main

import (
	"fmt"
	"os"

	"github.com/AgusDOLARD/zhare/internal"
	"github.com/AgusDOLARD/zhare/internal/server"
)

const (
	port = 3000
)

func main() {
	localIP, err := internal.GetLocalIP()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	url := fmt.Sprintf("http://%s:%d", localIP, port)
	err = internal.GenerateQR(url)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Printf("Serving on: %s", url)
	err = server.Serve(fmt.Sprintf(":%v", port), os.Args...)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func usage() {
	fmt.Printf("Usage: %s FILEPATH\n", os.Args[0])
}
