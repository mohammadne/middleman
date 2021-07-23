package request

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Post(url string, body interface{}) error {
	byteBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byteBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
