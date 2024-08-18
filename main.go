package main

import (
	"fmt"
	"log"
	"zhare/server"

	"github.com/alecthomas/kong"
)

var cli struct {
	Addr string `name:"addr" default:"localhost" help:"server address"`
	Port int    `name:"port" short:"p" default:"4000" help:"server port"`

	Files []string `arg:"" name:"file" type:"existingfile" help:"files to serve"`
}

func main() {
	kong.Parse(&cli)

	serverAddress := fmt.Sprintf("%s:%d", cli.Addr, cli.Port)
	srv := server.NewServer(serverAddress, cli.Files)
	err := srv.Start()
	if err != nil {
		log.Fatal(err)
	}
}
