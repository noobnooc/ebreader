package server

import (
	"fmt"
	"net/http"

	"github.com/hardo/ebreader/config"
)

//Run Listen the server
func Run() error {
	http.Handle("/", http.FileServer(http.Dir(config.Path)))
	address := fmt.Sprintf("%s:%d", config.Address, config.Port)
	fmt.Printf("Listening on %s ......\n", address)
	fmt.Printf("Please visit http://localhost:%d to view the book......\n", config.Port)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}
	return nil
}
