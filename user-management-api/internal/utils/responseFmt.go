package utils

type Response struct {
	Success bool         `json:"success"`
	Data    interface{}  `json:"data,omitempty"`
	Error   string       `json:"error,omitempty"`
	Meta    ResponseMeta `json:"meta,omitempty"`
}

type ResponseMeta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}