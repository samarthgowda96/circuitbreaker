package vinservice

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func NewVinService(id string) ([]byte, error) {
	url := fmt.Sprintf("https://cis-vin-decoder.p.rapidapi.com/vinDecode?vin=%s", id)
	apiKey := os.Getenv("RAPIDAPI_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("RAPIDAPI_KEY environment variable not set")
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "cis-vin-decoder.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
