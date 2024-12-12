package jobs

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func MakeGetRequest() (string, error) {
	url := "https://pingserver-lvkn.onrender.com"
	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error making GET request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	// Return the response as a string
	return string(body), nil
}
