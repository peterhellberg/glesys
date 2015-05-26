package glesys

import "net/url"

///////////////////////////////////////////////////////////////////////////////
// ip/listfree
///////////////////////////////////////////////////////////////////////////////

type IPListFreeResponse struct {
	Status Status `json:"status"`
	IPList struct {
		IPVersion  int    `json:"ipversion"`
		Datacenter string `json:"datacenter"`
		Platform   string `json:"platform"`
		CostPerIP  struct {
			Amount     int    `json:"amount"`
			Currency   string `json:"currency"`
			TimePeriod string `json:"timeperiod"`
		} `json:"costperip"`
		IPAddresses []string `json:"ipaddresses"`
	} `json:"iplist"`
	Debug Debug `json:"debug"`
}

// IPListFree returns a a list of all ip adresses that are available and not used on any account or server.
func (c *Client) IPListFree(ipversion, datacenter, platform string) (*IPListFreeResponse, error) {
	req, err := c.PostRequest("ip/listfree", url.Values{
		"ipversion":  {ipversion},
		"datacenter": {datacenter},
		"platform":   {platform},
	})
	if err != nil {
		return nil, err
	}

	var r struct {
		Response IPListFreeResponse `json:"response"`
	}

	resp, err := c.Do(req, &r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, ErrUnexpectedHTTPStatus
	}

	return &r.Response, nil
}
