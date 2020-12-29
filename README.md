# go-noaa
### Golang interface for NOAA API

<br/><br/>

# Getting Started

* Request your own token [here](https://www.ncdc.noaa.gov/cdo-web/token) to access the API.
* The token is provided as an environment variable to the package.
```bash
export TOKEN="your-token"
echo $TOKEN
> "your-token"
```
* For persistent storage of the environment variable, add `export TOKEN="your-token"` to `~/.bash_profile`.


<br/><br/>

## More Reading
* NOAA documentation on the API: https://www.ncdc.noaa.gov/cdo-web/webservices/v2
* Similar R package: https://github.com/ropensci/rnoaa