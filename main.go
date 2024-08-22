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
	Log     bool             `name:"log" default:"true" negatable:"" help:"disable logging"`
	Version kong.VersionFlag `short:"v" help:"Show version"`

	Files []string `arg:"" optional:"" name:"file" type:"existingfile" help:"files to serve"`
}

func main() {
	kong.Parse(&cli,
		kong.Vars{
			"version": version,
		})

	srv := server.NewServer(&server.ServerOpts{
		Addr:         fmt.Sprintf(":%d", cli.Port),
		Files:        cli.Files,
		EnableLogger: cli.Log,
	})
	err := srv.Start()
	if err != nil {
		log.Fatal(err)
	}
}
