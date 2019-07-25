package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("missing url! Please provide HAProxy url")
		os.Exit(1)
	}
	url := os.Args[1]

	var payloadSize int64 = 4 * 1024 * 1024 // 4 MB
	source := rand.New(rand.NewSource(20))
	limitedSource := io.LimitReader(source, payloadSize)

	req1, err := http.NewRequest(http.MethodPut, url, limitedSource)
	if err != nil {
		panic(err)
	}
	req1.ContentLength = payloadSize
	performRequest(http.DefaultClient, req1)

	fmt.Println()

	req2, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader([]byte(`{"message": "value"}`)))
	if err != nil {
		panic(err)
	}

	performRequest(http.DefaultClient, req2)
}

func performRequest(client *http.Client, req *http.Request) {
	fmt.Println("Sending request 1")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("Got response: %s\n", resp.Status)
	printHeaders(resp)
}

func printHeaders(resp *http.Response) {
	fmt.Println("Response headers:")
	for k, v := range resp.Header {
		fmt.Printf("%s: %s\n", k, v)
	}
}
