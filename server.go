package glesys

type ServerList struct {
	Response struct {
		Status  Status `json:"status"`
		Servers []struct {
			ServerID   string `json:"serverid"`
			Hostname   string `json:"hostname"`
			Datacenter string `json:"datacenter"`
			Platform   string `json:"platform"`
		}
		Debug Debug `json:"debug"`
	} `json:"response"`
}

func (c *Client) ServerList() (*ServerList, error) {
	req, err := c.GetRequest("server/list")
	if err != nil {
		return nil, err
	}

	var res ServerList

	if _, err = c.Do(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
