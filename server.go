package glesys

///////////////////////////////////////////////////////////////////////////////
// server/list
///////////////////////////////////////////////////////////////////////////////

// ServerListResponse contains the fields in a response from server/list
type ServerListResponse struct {
	Status  Status `json:"status"`
	Servers []struct {
		ServerID   string `json:"serverid"`
		Hostname   string `json:"hostname"`
		Datacenter string `json:"datacenter"`
		Platform   string `json:"platform"`
	}
	Debug Debug `json:"debug"`
}

// ServerList retrieves a list of servers on this account
func (c *Client) ServerList() (*ServerListResponse, error) {
	req, err := c.GetRequest("server/list")
	if err != nil {
		return nil, err
	}

	var r struct {
		Response ServerListResponse `json:"response"`
	}

	if _, err = c.Do(req, &r); err != nil {
		return nil, err
	}

	return &r.Response, nil
}
