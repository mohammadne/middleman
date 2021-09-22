package network

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Get(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	err = json.Unmarshal(bodyByte, body)
	if err != nil {
		return err
	}

	return nil
}
