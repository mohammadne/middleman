package handlers

import (
	"crypto/md5"

	"github.com/mohammadne/middleman/pkg/model"
	"github.com/mohammadne/middleman/pkg/request"
)

type ClientHandler struct {
	RequestUrl string
}

// GetObject will send get request to the url
func (ch ClientHandler) Get(key string) error {
	md5Key := md5.Sum([]byte(key))
	strKey := string(md5Key[:])
	body := model.Body{}

	return request.Get(ch.RequestUrl+"/"+strKey, &body)
}

// PostObject will send post request to the url
func (ch ClientHandler) Post(body interface{}) error {
	return request.Post(ch.RequestUrl, body)
}
