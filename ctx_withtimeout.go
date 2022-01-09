package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// TODO: set a http client timeout

	req, err := http.NewRequest("GET", "https://andcloud.io", nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(req.Context(), time.Millisecond*2)
	defer cancel() // to release resources if request does not timeout

	req = req.WithContext(ctx) //returns a shallow copy of req with its context changed to ctx

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	// Close the response body on the return.
	defer resp.Body.Close()

	// Write the response to stdout.
	io.Copy(os.Stdout, resp.Body)
}
