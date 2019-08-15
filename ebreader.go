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
	files.Clean()

	fmt.Println("Opening file......")
	err := files.Unepub()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Extracting file to working directory......")
	err = template.Build(path.Join(config.Path, "toc.ncx"))
	if err != nil {
		log.Fatalln(err)
	}

	err = server.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
