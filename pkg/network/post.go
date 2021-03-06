package network

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Post(url string, body interface{}) error {
	byteBody, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", "http://"+url, bytes.NewBuffer(byteBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
