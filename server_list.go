package glesys

import "time"

type ServerList struct {
	Response struct {
		Status struct {
			Code      int       `json:"code"`
			Timestamp time.Time `json:"timestamp"`
			Text      string    `json:"text"`
		} `json:"status"`
		Servers []struct {
			ServerID   string `json:"serverid"`
			Hostname   string `json:"hostname"`
			Datacenter string `json:"datacenter"`
			Platform   string `json:"platform"`
		}
	} `json:"response"`
}

func (c *Client) ServerList() (*ServerList, error) {
	req, err := c.NewRequest("GET", "server/list")
	if err != nil {
		return nil, err
	}

	var res ServerList

	if _, err = c.Do(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
