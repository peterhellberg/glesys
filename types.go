package glesys

import "time"

type Status struct {
	Code      int       `json:"code"`
	Timestamp time.Time `json:"timestamp"`
	Text      string    `json:"text"`
}

type Debug struct {
	Input map[string]interface{} `json:"input"`
}
