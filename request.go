package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func makeRequest(r string) {
	req, err := http.NewRequest("GET", r, nil) // Create a new http request
	req.Header.Add("token", token)             // Add authentication token
	client := &http.Client{}                   // Send request using http client
	resp, err := client.Do(req)                // Gather response
	handleError(err)                           // Handle error for response

	if resp.StatusCode != 200 {
		err := errors.New("Something went wrong")
		// Add some troubleshooting info here
		fmt.Println(resp.StatusCode)
		handleError(err)
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		handleError(err)
		d := make(map[string]interface{})
		/*
			Usually json data is unmarshalled into a
			struct with known fields. But here the response
			of the API is dynamic according to the endpoint.
			So, an empty interface type is used to store the
			unmarshalled json response.
		*/
		err = json.Unmarshal(body, &d) // Unmarshall and store in map
		handleError(err)
		jsonBytes, err := json.MarshalIndent(&d, "", " ") // Indentation if any
		handleError(err)
		ioutil.WriteFile("resp.json", jsonBytes, 0666) // Write to file
	}
}
