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

// SendHTTPSGet function to get from remote https URL
func SendHTTPSGet(url, bearerKey string) ([]byte, error) {

	res := []byte{}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", url, nil)
	if len(bearerKey) > 0 {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerKey))
	}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		res, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}
	} else {
		return res, fmt.Errorf("bad status code: %+v", resp.StatusCode)
	}
	return res, nil
}

// SendHTTPSRequest function to make a request to remote URL
func SendHTTPSRequest(url, bearerKey,
	method string, body []byte, contentType string) ([]byte, error) {

	res := []byte{}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	if len(bearerKey) > 0 {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerKey))
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
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		res, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}
	} else {
		return res, fmt.Errorf("bad status code: %+v", resp.StatusCode)
	}
	return res, nil
}

// SendHTTPSFileRequest function to POST file from provided PATH
func SendHTTPSFileRequest(url, bearerKey string, path string) ([]byte, error) {
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
	if err != nil {
		return res, err
	}
	err = writer.Close()
	if err != nil {
		return res, err
	}

	return SendHTTPSRequest(url, bearerKey, "POST", body.Bytes(), writer.FormDataContentType())
}

// SendHTTPSFileDataRequest function to POST file & data
func SendHTTPSFileDataRequest(url, bearerKey string, path string, data []byte) ([]byte, error) {
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
	if err != nil {
		return res, err
	}
	err = writer.Close()
	if err != nil {
		return res, err
	}

	return SendHTTPSRequest(url, bearerKey, "POST", body.Bytes(), writer.FormDataContentType())
}
