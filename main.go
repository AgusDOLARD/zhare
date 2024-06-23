package main

import (
	"flag"
	"io"
	"log/slog"
	"os"
	"strconv"
	"zhare/server"
)

var (
	portFlag  int
	debugFlag bool
	reader    io.Reader
)

func main() {
	flag.IntVar(&portFlag, "p", 3000, "server port")
	flag.BoolVar(&debugFlag, "debug", false, "debug mode")
	flag.Parse()

	if !debugFlag {
		slog.SetLogLoggerLevel(slog.LevelWarn)
	}

	if len(flag.Args()) == 0 {
		reader = os.Stdin
	} else {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			slog.Error("Error opening file", "err", err.Error())
			os.Exit(1)
		}
		reader = f
	}

	srv := server.NewServer(":"+strconv.Itoa(portFlag), reader)
	if err := srv.Start(); err != nil {
		slog.Error("Server Error", "err", err)
	}
}
