package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	w := webData{
		url:    "https://www.mwjcomputing.com",
		method: "GET",
	}

	printHeader(w)
	getHeaders(w)
}

type webData struct {
	url, method string
}

func printHeader(data webData) {
	fmt.Println("PoshGoWebHeaders")
	fmt.Println("================")
	fmt.Printf("URL: %s\n", data.url)
	fmt.Printf("HTTP Method: %s\n", data.method)
	fmt.Println("================")
}

// getHeaders retreives the header values of a supplied website and HTTP Method.
// It displays the content of the headers to the screen, and returns nil.
func getHeaders(data webData) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// req variable holds an HTTP request object. It takes the HTTP Method, and a URL as its parameters.
	req, err := http.NewRequest(data.method, data.url, nil)
	if err != nil {
		return fmt.Errorf("Error received %s", err.Error())
	}

	// res variable holds an HTTP response object. It takes the request object.
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error received %s", err.Error())
	}

	// This closes the response stream.
	defer res.Body.Close()

	contentTypesValues := res.Header.Get("content-type")
	fmt.Printf("Content Type: %s\n", contentTypesValues)

	// valXFrameOptions gets the X-Frame-Options header from the web request.
	valXFrameOptions := res.Header.Get("X-Frame-Options")
	if "" != valXFrameOptions {
		fmt.Printf("X-Frame-Options: %s\n", valXFrameOptions)
	}

	return nil
}
