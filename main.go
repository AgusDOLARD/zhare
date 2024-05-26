package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AgusDOLARD/zhare/internal"
)

var (
	qrFlag   bool
	portFlag int
)

func main() {
	flag.BoolVar(&qrFlag, "qr", false, "show qr for web page")
	flag.IntVar(&portFlag, "p", 3000, "server port")
	flag.Parse()

	if len(flag.Args()) == 0 {
		usage()
		os.Exit(1)
	}

	localIP, err := internal.GetLocalIP()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	url := fmt.Sprintf("http://%s:%d", localIP, portFlag)

	if qrFlag {
		err = internal.GenerateQR(url)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}
	}

	fmt.Printf("Serving on: %s", url)
	err = internal.Serve(fmt.Sprintf(":%v", portFlag), flag.Arg(0))
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func usage() {
	fmt.Printf("Usage: %s DIRECTORY [OPTIONS]\n", os.Args[0])
}
