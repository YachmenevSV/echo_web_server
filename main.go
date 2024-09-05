package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func getRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("got / request\n")
	b, _ := httputil.DumpRequest(req, false)
	io.WriteString(w, string(b))

	b, err := io.ReadAll(req.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Wait while you press enter button\n")
	fmt.Scanln()
	io.WriteString(w, "\nYour body:\n")
	io.WriteString(w, string(b))
}

func main() {
	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":41000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
