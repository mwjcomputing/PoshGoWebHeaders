package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	urlPtr := flag.String("url", "https://www.mwjcomputing.com", "URL to check")
	methodPtr := flag.String("method", "https://www.mwjcomputing.com", "HTTP Method to use")
	flag.Parse()

	w := webData{
		url:    *urlPtr,
		method: *methodPtr,
	}

	printHeader(w)

	err := getHeaders(w)
	if err != nil {
		return
	}
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

	return
}

// getHeaders retrieves the header values of a supplied website and HTTP Method.
// It displays the content of the headers to the screen, and returns nil.
func getHeaders(data webData) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// req variable holds an HTTP request object. It takes the HTTP Method, and a URL as its parameters.
	req, err := http.NewRequest(data.method, data.url, nil)
	if err != nil {
		return fmt.Errorf("error received %s", err.Error())
	}

	// res variable holds an HTTP response object. It takes the request object.
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error received %s", err.Error())
	}

	// This closes the response stream.
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Cannot close the reader. Error: %s", err.Error())
		}
	}(res.Body)

	var (
		valNotFound = "Value not found"
	)

	contentTypesValues := res.Header.Get("content-type")
	fmt.Printf("Content Type: %s\n", contentTypesValues)

	// valXFrameOptions gets the X-Frame-Options header from the web request.
	valXFrameOptions := res.Header.Get("X-Frame-Options")
	if "" != valXFrameOptions {
		fmt.Printf("X-Frame-Options: %s\n", valXFrameOptions)
	} else {
		fmt.Printf("X-Frame-Options: %s\n", valNotFound)
	}

	valXSSProtection := res.Header.Get("X-XSS-Protection")
	if "" != valXSSProtection {
		fmt.Printf("X-XSS-Protection: %s\n", valXSSProtection)
	} else {
		fmt.Printf("X-XSS-Protection: %s\n", valNotFound)
	}
	return nil
}
