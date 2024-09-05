package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
)

func getRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("got / request\n")

	// Печать сообщения и ожидание ввода от пользователя
	fmt.Printf("Enter the response status code or press Enter to use 200: ")
	var userInput string
	fmt.Scanln(&userInput)

	// Если пользователь ввел код ответа, преобразуем его в int
	statusCode := 200 // По умолчанию 200
	if userInput != "" {
		if code, err := strconv.Atoi(userInput); err == nil {
			statusCode = code
		} else {
			fmt.Printf("Invalid input, using default 200\n")
		}
	}

	w.WriteHeader(statusCode)
	fmt.Printf("Set response code: %d\n", statusCode)
	b, _ := httputil.DumpRequest(req, false)
	io.WriteString(w, string(b))

	b, err := io.ReadAll(req.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}
	// Устанавливаем код ответа

	io.WriteString(w, "\nYour body:\n")
	io.WriteString(w, string(b))
}

func main() {
	fmt.Printf("Starting listen on localhost:41000 ")

	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":41000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
