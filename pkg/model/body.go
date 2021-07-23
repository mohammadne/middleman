package model

type Body struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Cache bool   `json:"cache"`
}
