package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func authenticate() {

	checkEnv(token)            // Make sure environment variable is set
	token = os.Getenv("TOKEN") // Get token

	// Generate a test url
	endpoint := url + "data?datasetid=GHCND&locationid=ZIP:28801&startdate=2010-05-01&enddate=2010-05-01"
	req, err := http.NewRequest("GET", endpoint, nil) // Create a new http request
	req.Header.Add("token", token)                    // Add authentication token
	client := &http.Client{}                          // Send request using http client
	resp, err := client.Do(req)                       // Gather response
	handleError(err)                                  // Handle error for response
	// 2nd line of defense: Check response code 200
	if resp.StatusCode != 200 {
		err := errors.New("Something went wrong")
		fmt.Println(resp.StatusCode)
		handleError(err)
	} else {
		fmt.Println("Response OK. Token authenticated")
		return
	}
}

func checkEnv(token string) {
	_, exists := os.LookupEnv("TOKEN")
	if !exists {
		err := errors.New("$TOKEN environment variable has not been set")
		handleError(err)
	}
}
