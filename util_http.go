package libwimark

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func SendHTTPPost(url, mime string, request []byte) ([]byte, error) {
	if len(mime) == 0 {
		mime = "application/json"
	}
	resp, err := http.Post(url, mime, bytes.NewBuffer(request))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func SendHTTPGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func SendHTTPPostFile(url, filename, filetype string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(filetype, filepath.Base(file.Name()))
	if err != nil {
		return []byte{}, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return []byte{}, err
	}
	writer.Close()
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return []byte{}, err
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
