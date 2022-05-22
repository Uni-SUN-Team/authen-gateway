package models

type Response struct {
	Error  string            `json:"error"`
	Result map[string]string `json:"result"`
}
