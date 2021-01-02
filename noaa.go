/*
Author: Harry Singh
Summary:

*/

package main

import (
	"flag"
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

	/*
		User can either setup config.yaml and pass it to the program
		Or run it without the config file and entering parameters
		during runtime with user step-wise prompts.
	*/
	var configFile string
	flag.StringVar(&configFile, "config", "", "Pass config file name (config.yaml)")

	flag.Parse()

	var config map[string]string
	if configFile == "" {
		config = buildConfig() // Build config with prompts if no file passed
	} else {
		config = loadConfig(configFile) // Load from file if passed with flag
	}

	/*
		Config is a map object with key:value pairs corresponding to yaml fields
		The first key:value is the endpoint
		Rest are the optional params
	*/

	r := buildRequest(config)

	fmt.Println(r)
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return
}
