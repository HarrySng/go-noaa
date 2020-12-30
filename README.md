# go-noaa
### Golang interface for NOAA API: https://www.ncdc.noaa.gov/cdo-web/webservices/v2
<br/>

## Note: Package Under Development
Check back soon for updates

<br/>

# Getting Started

## Prerequisites
* Download the binary from binaries/ according to your OS.
* Binaries are stand-alone. They have no prerequisites except the OS platform.

<br/>

## Token
* Request your own token [here](https://www.ncdc.noaa.gov/cdo-web/token) to access the API.
* The token is provided as an environment variable to the package.
```bash
export TOKEN="your-token"
echo $TOKEN
> "your-token"
```
* For persistent storage of the environment variable, add `export TOKEN="your-token"` to `~/.bash_profile`.
<br/>

## Configuration file
* There are two methods of configuring the web request.
1. Set up the `config.yaml` file and pass it to the program with `-config config.yaml` flag.
2. Run the program without the flag and select prompts while the program executes.

### Method 1: Setting up `config.yaml`
* The file is divided into blocks corresponding to each endpoint of the API.
* Each block has `key:value` pairs corresponding to parameters passed to web request.
* Uncomment only one block and fill in the optional parameters as required.
* Save the file and run the program with `go-noaa -config config.yaml`


### Method 2: Run the program without config.yaml
* Run the program without the config flag `go-noaa`
* Follow the prompts on the command line and select the parameters.


<br/><br/>

