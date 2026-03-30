package functions

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type FullResponse struct {
	Location    string
	ContentType string
	Body        struct {
		Input struct {
			Size int    `json:"size"`
			Type string `json:"type"`
		} `json:"input"`
	}
}

func image_resize(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.tinify.com/shrink", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	req.SetBasicAuth("api", API)
	req.Header.Set("Host", "api.tinify.com")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var full FullResponse
	full.Location = resp.Header.Get("Location")
	full.ContentType = resp.Header.Get("Content-Type")
	json.NewDecoder(resp.Body).Decode(&full.Body)

}
