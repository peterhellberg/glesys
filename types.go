package glesys

import "time"

type ResponseStatus struct {
	Code      int       `json:"code"`
	Timestamp time.Time `json:"timestamp"`
	Text      string    `json:"text"`
}

type ResponseDebug struct {
	Input map[string]interface{} `json:"input"`
}
