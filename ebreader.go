package main

import (
	"ebreader/config"
	"ebreader/server"
	"ebreader/util/files"
	"ebreader/util/template"
	"fmt"
	"log"
	"path"
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
