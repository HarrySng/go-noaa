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

func main() {

	authenticate() // Authentication checkpoint to verify token
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return
}
