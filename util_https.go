package libwimark

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func SendHTTPSGet(url, apiKey string) ([]byte, error) {

	res := []byte{}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", url, nil)
	if len(apiKey) > 0 {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return res, err
	}

	if resp.StatusCode == http.StatusOK {
		res, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}
	} else {
		return res, fmt.Errorf("Bad status code: %+v", resp.StatusCode)
	}
	return res, nil
}

func SendHTTPSRequest(url, apiKey string,
	method string, body []byte, contentType string) ([]byte, error) {

	res := []byte{}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	if len(apiKey) > 0 {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	}
	switch contentType {
	case "plain":
		req.Header.Set("Content-Type", "text/plain")
	case "", "json":
		req.Header.Set("Content-Type", "application/json")
	default:
		req.Header.Set("Content-Type", contentType)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return res, err
	}

	if resp.StatusCode == http.StatusOK {
		res, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}
	} else {
		return res, fmt.Errorf("Bad status code: %+v", resp.StatusCode)
	}
	return res, nil
}

func SendHTTPSFileRequest(url, apiKey string, path string) ([]byte, error) {
	res := []byte{}
	file, err := os.Open(path)
	if err != nil {
		return res, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return res, err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return res, err
	}

	return SendHTTPSRequest(url, apiKey, "POST", body.Bytes(), writer.FormDataContentType())
}

func SendHTTPSFileDataRequest(url, apiKey string, path string, data []byte) ([]byte, error) {
	res := []byte{}
	file, err := os.Open(path)
	if err != nil {
		return res, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if len(data) > 0 {
		err = writer.WriteField("data", string(data))
		if err != nil {
			return res, err
		}
	}

	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return res, err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return res, err
	}

	return SendHTTPSRequest(url, apiKey, "POST", body.Bytes(), writer.FormDataContentType())
}
