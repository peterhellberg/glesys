package glesys

import "net/url"

type IPListFree struct {
	Response struct {
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
	} `json:"response"`
}

func (c *Client) IPListFree(ipversion, datacenter, platform string) (*IPListFree, error) {
	req, err := c.PostRequest("ip/listfree", url.Values{
		"ipversion":  {ipversion},
		"datacenter": {datacenter},
		"platform":   {platform},
	})
	if err != nil {
		return nil, err
	}

	var listfree IPListFree

	resp, err := c.Do(req, &listfree)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return &listfree, ErrUnexpectedHTTPStatus
	}

	return &listfree, nil
}
