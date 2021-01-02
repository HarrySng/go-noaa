package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"gopkg.in/yaml.v2"
)

func buildConfig() map[string]string {

	// Slice of endpoints
	e := []string{"datasets", "datacategories", "datatypes", "locationcategories", "locations", "stations", "data"}

	// Map each endpoint to an int for user input
	endpoints := make(map[int]string)
	for i := 0; i < 7; i++ {
		endpoints[i] = e[i]
	}

	/*
		Map object does not store items in order of ints
		To force it, create a separate int slice
	*/
	keys := make([]int, len(endpoints))
	j := 0
	for k := range endpoints {
		keys[j] = k
		j++
	}
	sort.Ints(keys) // Sort the int slice

	/*
		Then print the slice and corresponding map value
		rather than printing the map itself.
	*/
	fmt.Println("\nChoose an endpoint:")
	for _, k := range keys {
		fmt.Println(k, endpoints[k])
	}

	fmt.Println("\nEnter a number from 0 to 6:")
	var i int
	_, err := fmt.Scanf("%d", &i) // Ask user for input
	handleError(err)

	/*
		Start building the config object now.
		Goal is to build the object exactly similar
		to if a user passed the config file.
		Means the return object from both functions
		will be exactly similar.
	*/
	config := make(map[string]string)

	if i >= 0 && i < 7 {
		config["endpoint"] = endpoints[i] // Insert the endpoint in config
		// Prompt for common parameters first, then handle edge cases

		/*
			Sanity checks are not done on each user input
			as it is too tedious.
			In case user inputs incorrect parameter value,
			the request will fail.
		*/
		config["startdate"] = ask("\nEnter start date (1970-10-03) or leave blank: ")
		config["enddate"] = ask("\nEnter end date (2010-09-10) or leave blank: ")
		config["sortfield"] = ask("\nEnter field to sort results by (name): ")

		sortorder := ask("\nType 'desc' to change sort order or press enter for 'asc':")
		if sortorder == "" {
			sortorder = "asc"
		}
		config["sortorder"] = sortorder

		limit := ask("\nSet limit of results returned or press Enter for default (25): ")
		if limit == "" {
			limit = "25"
		}
		config["limit"] = limit

		offset := ask("\nEnter offset value or press Enter for default (0): ")
		if offset == "" {
			offset = "0"
		}
		config["offset"] = offset

		// Now handle individual parameters case by case
		// Some repetition here. Will remove later

		if i == 0 { // datasets
			datatypeid := ask("\nEnter a datatypeid (ACMH) or press Enter to leave blank: ")
			config["datatypeid"] = datatypeid

			locationid := ask("\nEnter a locationid (FIPS:37) or press Enter to leave blank: ")
			config["locationid"] = locationid

			stationid := ask("\nEnter a stationid (COOP:010957) or press Enter to leave blank: ")
			config["stationid"] = stationid

		} else if i == 1 { // datacategories
			datasetid := ask("\nEnter a datasetid (GCOM) or press Enter to leave blank: ")
			config["datasetid"] = datasetid

			locationid := ask("\nEnter a locationid (FIPS:37) or press Enter to leave blank: ")
			config["locationid"] = locationid

			stationid := ask("\nEnter a stationid (COOP:010957) or press Enter to leave blank: ")
			config["stationid"] = stationid

		} else if i == 2 { // datatypes
			datasetid := ask("\nEnter a datasetid (GCOM) or press Enter to leave blank: ")
			config["datasetid"] = datasetid

			locationid := ask("\nEnter a locationid (FIPS:37) or press Enter to leave blank: ")
			config["locationid"] = locationid

			stationid := ask("\nEnter a stationid (COOP:010957) or press Enter to leave blank: ")
			config["stationid"] = stationid

			datacategoryid := ask("\nEnter a datacategoryid (TEMP) or press Enter to leave blank: ")
			config["datacategoryid"] = datacategoryid

		} else if i == 3 { // locationcategories
			datasetid := ask("\nEnter a datasetid (GCOM) or press Enter to leave blank: ")
			config["datasetid"] = datasetid

		} else if i == 4 { // locations
			datasetid := ask("\nEnter a datasetid (GCOM) or press Enter to leave blank: ")
			config["datasetid"] = datasetid

			locationcategoryid := ask("\nEnter a locationcategoryid (CITY) or press Enter to leave blank: ")
			config["locationcategoryid"] = locationcategoryid

			datacategoryid := ask("\nEnter a datacategoryid (TEMP) or press Enter to leave blank: ")
			config["datacategoryid"] = datacategoryid

		} else if i == 5 { // stations
			datasetid := ask("\nEnter a datasetid (GCOM) or press Enter to leave blank: ")
			config["datasetid"] = datasetid

			locationid := ask("\nEnter a locationid (FIPS:37) or press Enter to leave blank: ")
			config["locationid"] = locationid

			datatypeid := ask("\nEnter a datatypeid (ACMH) or press Enter to leave blank: ")
			config["datatypeid"] = datatypeid

			datacategoryid := ask("\nEnter a datacategoryid (TEMP) or press Enter to leave blank: ")
			config["datacategoryid"] = datacategoryid

			extent1 := ask("\nEnter minimum latitude (47.50): ")
			config["extent1"] = extent1
			extent2 := ask("\nEnter minimum longitude (-122.50): ")
			config["extent2"] = extent2
			extent3 := ask("\nEnter maximum latitude (48.50): ")
			config["extent3"] = extent3
			extent4 := ask("\nEnter maximum longitude (-121.50): ")
			config["extent4"] = extent4

		} else { // data
			datasetid := ask("\nEnter a datasetid (GCOM) or press Enter to leave blank: ")
			config["datasetid"] = datasetid

			datatypeid := ask("\nEnter a datatypeid (ACMH) or press Enter to leave blank: ")
			config["datatypeid"] = datatypeid

			locationid := ask("\nEnter a locationid (FIPS:37) or press Enter to leave blank: ")
			config["locationid"] = locationid

			stationid := ask("\nEnter a stationid (COOP:010957) or press Enter to leave blank: ")
			config["stationid"] = stationid

			units := ask("\nEnter units (metric): ")
			config["units"] = units

			includemetadata := ask("\nEnter units (metric): ")
			if includemetadata == "" {
				includemetadata = "true"
			}
			config["includemetadata"] = includemetadata
		}
	} else {
		err := errors.New("You did not enter a valid option. Try again")
		handleError(err)
	}

	return config
}

func ask(q string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(q)
	s, err := reader.ReadString('\n')
	handleError(err)
	if s == "\n" { // If user skips, it returns a new line
		s = "" // Replace with blank string
	}
	return s
}

func loadConfig(configFile string) map[string]string {
	c := cData{}                              // Initiate the struct
	yfile, err := ioutil.ReadFile(configFile) // Read file
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
