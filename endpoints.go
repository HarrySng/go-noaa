package main

/*
To fix: All optional parameters are not appended
with ?. Only the first is, rest are appended with &
*/

// Endpoint is an interface to hold config.yaml0
type Endpoint map[string]string

// cData is a struct to hold unmarshalled config.yaml data
type cData struct {
	Endpoint []Endpoint `yaml:"endpoints"`
}

func buildRequest(config map[string]string) string {
	var r string

	// Now build r step by step

	// Add the endpoint first
	r = url + config["endpoint"]

	/*
		There are several cases here.
		- endpoint key needs to be checked within the
		loop as well as order of keys is random.
		- Different format
			- url/datasets
			- url/datasets/GSOY
			- url/datasets?datatypeid=TOBS
			- url/datatypes?datacategoryid=TEMP&limit=56
			- url/datasets/GSOY&ACMH
	*/

	// First handling the IDs which are appended with forward slash
	// All these use datasetid so use for loop
	var rid string // Initiate a separate string url
	d := []string{"datacategories", "datatypes", "locationcategories", "locations", "stations"}
	for i := range d {
		rid = appendID(config, r, d[i], "datasetid") // Append datasetis IF it exists
		if r != rid {
			return rid // Dont go to next iteration if rid changed which means ID was provided the endpoint being iterated
		}
	}
	rid = appendID(config, r, "datasets", "datatypeid") // This uses datatypeid so separate

	// If ID provided, do not check rest of the parameters
	if r != rid { // ID was provided
		return rid // Return rid and stop further actions
	}

	// If ID not provided, continue building from r (not rid)
	// All the rest are appended with ?par=
	r = appendOptPars(config, r)

	return r
}

func appendID(config map[string]string, r string, v string, k string) string {
	/*
		The function checks that whether an ID (k) was passed
		to the particular endpoint being called (v).
	*/
	if config["endpoint"] == v { // For the endpoint
		if config[k] != "" { // Check if an ID was passed (different IDs for different endpoints)
			r = r + "/" + config[k] // If yes then append it after a forward slash
		}
	}
	return r
}

func appendOptPars(config map[string]string, r string) string {

	/*
		Handle steps for each case
		There are several overlaps within endpoints
		So they are grouped together with ||
		Then pending cases are handled individually or through
		grouping with another common endpoint.

		r is defined in the parent function and passed
		to this function where it is built successively and returned.

		Extent case to be added later for stations/
	*/

	if config["endpoint"] == "datasets" || config["endpoint"] == "datacategories" || config["endpoint"] == "datatypes" {
		// locationid and stationid common in these 3 endpoints
		if config["locationid"] != "" {
			r = r + "?" + "locationid" + "=" + config["locationid"]
		}
		if config["stationid"] != "" {
			r = r + "?" + "stationid" + "=" + config["stationid"]
		}
	}

	// datacategoryid common in these endpoints
	if config["endpoint"] == "datatypes" || config["endpoint"] == "locations" || config["endpoint"] == "stations" {
		if config["datacategoryid"] != "" {
			r = r + "?" + "datacategoryid" + "=" + config["datacategoryid"]
		}
	}

	if config["endpoint"] == "locations" {
		if config["locationcategoryid"] != "" {
			r = r + "?" + "locationcategoryid" + "=" + config["locationcategoryid"]
		}
	}

	if config["endpoint"] == "stations" || config["endpoint"] == "data" {
		// locationid and datatypeid common in these two endpoints
		if config["locationid"] != "" {
			r = r + "?" + "locationid" + "=" + config["locationid"]
		}
		if config["datatypeid"] != "" {
			r = r + "?" + "datatypeid" + "=" + config["datatypeid"]
		}
	}

	// Separate case for data
	if config["endpoint"] == "data" {
		if config["datasetid"] != "" {
			r = r + "?" + "datasetid" + "=" + config["datasetid"]
		}
		if config["stationid"] != "" {
			r = r + "?" + "stationid" + "=" + config["stationid"]
		}

		d := []string{"units", "includemetadata"}
		for i := range d {
			if config[d[i]] != "" {
				r = r + "?" + d[i] + "=" + config[d[i]]
			}
		}
	}

	//  Append all the common optional parameters
	d := []string{"startdate", "enddate", "sortfield", "sortorder", "limit", "offset"}
	for i := range d {
		if config[d[i]] != "" {
			r = r + "?" + d[i] + "=" + config[d[i]]
		}
	}
	return r
}
