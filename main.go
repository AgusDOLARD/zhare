package main

import (
	"fmt"
	"log"
	"zhare/server"

	"github.com/alecthomas/kong"
)

var version = "dev"

var cli struct {
	Port    int              `name:"port" short:"p" default:"4000" help:"server port"`
	Version kong.VersionFlag `short:"v" help:"Show version"`

	Files []string `arg:"" name:"file" type:"existingfile" help:"files to serve"`
}

func main() {
	kong.Parse(&cli,
		kong.Vars{
			"version": version,
		})

	serverAddress := fmt.Sprintf(":%d", cli.Port)
	srv := server.NewServer(serverAddress, cli.Files)
	err := srv.Start()
	if err != nil {
		log.Fatal(err)
	}
}
