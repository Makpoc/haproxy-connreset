package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("missing port")
		os.Exit(1)
	}
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to parse port: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Starting on port: :%d\n", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			_, err = io.CopyN(ioutil.Discard, r.Body, 1024)
			if err != nil {
				fmt.Println(err)
			}
			w.WriteHeader(http.StatusNoContent)
			return
		}
		_, _ = io.Copy(ioutil.Discard, r.Body)
		w.WriteHeader(http.StatusOK)
	}))
	log.Fatal(err)
}
