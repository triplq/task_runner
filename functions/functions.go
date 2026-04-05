package functions

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type BodyRequest struct {
	Resize struct {
		Method string `json:"method"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"resize"`
}

func Image_resize(path string, w, h int) error {
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

	loc := resp.Header.Get("Location")

	var resize BodyRequest
	resize.Resize.Method = "fit"
	resize.Resize.Width = w
	resize.Resize.Height = h

	bodyBytes, err := json.Marshal(resize)
	if err != nil {
		return err
	}

	req, err = http.NewRequest("POST", loc, bytes.NewBuffer(bodyBytes))
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

	outpath := filepath.Dir(path)
	outfile := filepath.Base(path)
	filename := strings.Split(outfile, ".")

	fp := filepath.Join(outpath, filename[0]+"_thumb."+filename[1])

	out, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}
