package main

import (
	"fmt"
	"log"
	"os"
	"zhare/server"

	"github.com/alecthomas/kong"
)

var version = "dev"

var cli struct {
	Port      int              `name:"port" short:"p" default:"4000" help:"server port"`
	Directory string           `name:"dir" short:"d" type:"existingdir" help:"directory to serve"`
	Log       bool             `name:"log" default:"true" negatable:"" help:"disable logging"`
	Version   kong.VersionFlag `short:"v" help:"Show version"`

	Files []string `arg:"" optional:"" name:"file" type:"existingfile" help:"files to serve"`
}

func main() {
	kong.Parse(&cli,
		kong.Vars{
			"version": version,
		})

	if cli.Directory != "" {
		cli.Files = getFilesFromDir(cli.Directory)
	}

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

func getFilesFromDir(dir string) []string {
	fileNames := make([]string, 0)
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			fileNames = append(fileNames, f.Name())
		}
	}
	return fileNames
}
