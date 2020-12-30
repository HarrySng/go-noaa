package main

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func loadConfig() map[string]string {
	f := "config.yaml"               // Config file
	c := cData{}                     // Initiate the struct
	yfile, err := ioutil.ReadFile(f) // Read file
	handleError(err)
	err = yaml.Unmarshal(yfile, &c) // Unmarshall
	handleError(err)

	config := make(map[string]string) // Create an empty map to hold configs

	for _, v := range c.Endpoint { // Iterate over parent element (- endpoints)
		// Ignore its key as there is no key, the value holds each config block
		for k, vv := range v { // Within each config block, extract key, value pair
			config[k] = vv // Push values into empty map object
		}
	}
	/*
		Sanity check on the config that may have been overlooked by Unmarshal.
		1. Check that the mandatory key "endpoint" is present.
		2. Check that there are no duplicate keys uncommented in the file.
	*/
	if _, ok := config["endpoint"]; !ok {
		err := errors.New("The mandatory 'endpoint' configuration key is missing. Please check config.yaml")
		handleError(err)
	}
	return config
}
