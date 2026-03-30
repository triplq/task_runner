package functions

import (
	"bytes"
	"encoding/json"
	"io"
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

type BodyRequest struct {
	Resize struct {
		Method string `json:"method"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"resize"`
}

func image_resize(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.tinify.com/shrink", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.SetBasicAuth("api", API)
	req.Header.Set("Host", "api.tinify.com")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var full FullResponse
	full.Location = resp.Header.Get("Location")
	full.ContentType = resp.Header.Get("Content-Type")
	json.NewDecoder(resp.Body).Decode(&full.Body)

	var resize BodyRequest
	resize.Resize.Method = "fit"
	resize.Resize.Width = 600
	resize.Resize.Height = 600

	bodyBytes, err := json.Marshal(resize)
	if err != nil {
		return err
	}

	req, err = http.NewRequest("POST", full.Location, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}

	req.SetBasicAuth("api", API)
	req.Header.Set("Host", "api.tinify.com")
	req.Header.Set("Content-Type", "application/json")

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create("../data/panda_thumb.jpg")
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}
