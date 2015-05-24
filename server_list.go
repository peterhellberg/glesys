package glesys

import (
	"encoding/json"
	"time"
)

type ResponseStatus struct {
	Code      int       `json:"code"`
	Timestamp time.Time `json:"timestamp"`
	Text      string    `json:"text"`
}

type ResponseDebug struct {
	Input json.RawMessage `json:"input"`
}

type ServerList struct {
	Response struct {
		Status  ResponseStatus `json:"status"`
		Servers []struct {
			ServerID   string `json:"serverid"`
			Hostname   string `json:"hostname"`
			Datacenter string `json:"datacenter"`
			Platform   string `json:"platform"`
		}
		Debug ResponseDebug `json:"debug"`
	} `json:"response"`
}

func (c *Client) ServerList() (*ServerList, error) {
	req, err := c.NewRequest("GET", "server/list", nil)
	if err != nil {
		return nil, err
	}

	var res ServerList

	if _, err = c.Do(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
