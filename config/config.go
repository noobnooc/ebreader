package config

import (
	"flag"
	"fmt"
	"os"
	"path"
)

const defaultPort = 4444

var (
	//File The epub file to open
	File string
	//Port The port to listen
	Port int
	//Address The address to listen
	Address string
	//Path Working directory
	Path string
)

func init() {
	initOptions()
	Path += path.Join(Path, "ebreader")
}

func initOptions() {
	flag.IntVar(&Port, "p", defaultPort, "The port ebreader will listen on.")
	flag.StringVar(&Path, "w", os.TempDir(), "The working directory of ebreader.")
	flag.StringVar(&Address, "a", "127.0.0.1", "The address to listen.")
	flag.Parse()

	if flag.NArg() != 1 {
		showHelper()
	}
	File = flag.Arg(0)
}

func showHelper() {
	fmt.Println("Usage：")
	fmt.Println("  ebreader [-options=value] filename")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println()
	flag.PrintDefaults()
	os.Exit(2)
}
