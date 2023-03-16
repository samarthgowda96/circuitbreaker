package vinservice

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

// Fetch VIN details from the RapidAPI
func NewVinService(id string) ([]byte, error) {
	// get the RapidAPI key from the environment
	apiKey, ok := os.LookupEnv("RAPIDAPI_KEY")
	if !ok {
		return nil, errors.New("RAPIDAPI_KEY environment variable not set")
	}

	// create a new HTTP request with the VIN ID and RapidAPI key
	url := "https://cis-vin-decoder.p.rapidapi.com/vinDecode?vin=" + id
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("X-RapidAPI-Key", apiKey)

	// define a function to execute the HTTP request
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to fetch data")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}
