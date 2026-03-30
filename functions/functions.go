package functions

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func image_resize(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.tinify.com/shrink", bytes.NewBuffer(data))

	req.SetBasicAuth("api", API)
	req.Header.Set("Host", "api.tinify.com")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	return string(body), nil

}
