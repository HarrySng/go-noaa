package main

import "fmt"

func buildRequest(config map[string]string) string {
	var r string

	// Now build r step by step

	// Adding outside loop as this is minimum requirement
	r = url + config["endpoint"]

	for k, v := range config {
		/*
			There are several cases here.
			- endpoint key needs to be checked within the
			loop as well as order of keys is random.
			- Different format
				- url/datasets
				- url/datasets/GSOY
				- url/datasets?datatypeid=TOBS
				- url/datatypes?datacategoryid=TEMP&limit=56
		*/

		if k == "endpoint" {
			continue
		} else {
			fmt.Println(k, v)
		}

	}
	return r
}
