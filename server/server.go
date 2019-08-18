package server

import (
	"fmt"
	"net/http"
)

//Run Listen the server
func Run(address string, port int, path string) error {
	http.Handle("/", http.FileServer(http.Dir(path)))
	a := fmt.Sprintf("%s:%d", address, port)
	fmt.Printf("Listening on %s ......\n", a)
	fmt.Printf("Please visit http://localhost:%d to view the book......\n", port)
	err := http.ListenAndServe(a, nil)
	if err != nil {
		return err
	}
	return nil
}
