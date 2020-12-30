/*
Author: Harry Singh
Summary:

*/

package main

import (
	"fmt"
	"os"
)

var url = "https://www.ncdc.noaa.gov/cdo-web/api/v2/"
var token string

// Endpoint is an interface to hold config.yaml0
type Endpoint map[string]string

// cData is a struct to hold unmarshalled config.yaml data
type cData struct {
	Endpoint []Endpoint `yaml:"endpoints"`
}

func main() {

	authenticate() // Authentication checkpoint to verify token
	config := loadConfig()
	/*
		Load and parse config.yaml
		Config is a map object with key:value pairs corresponding to yaml fields
		The first key:value is the endpoint
		Rest are the optional params
	*/
	fmt.Println(config)
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return
}
