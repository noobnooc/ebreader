package main

import (
	"fmt"
	"log"
	"path"

	"github.com/hardo/ebreader/config"
	"github.com/hardo/ebreader/server"
	"github.com/hardo/ebreader/util/files"
	"github.com/hardo/ebreader/util/template"
)

func main() {
	start()
}

func start() {
	files.Clean(config.Path)

	fmt.Println("Opening file......")
	err := files.Unepub(config.Path, config.File)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Extracting file to working directory......")
	err = template.Build(config.Path, path.Join(config.Path, "toc.ncx"))
	if err != nil {
		log.Fatalln(err)
	}

	err = server.Run(config.Address, config.Port, config.Path)
	if err != nil {
		log.Fatalln(err)
	}
}
