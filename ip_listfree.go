package glesys

import (
	"fmt"
	"net/url"
	"strings"
)

type IPListFree struct {
	Response struct {
		Status ResponseStatus `json:"status"`
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
		Debug ResponseDebug `json:"debug"`
	} `json:"response"`
}

func (c *Client) IPListFree(ipversion, datacenter, platform string) (*IPListFree, error) {
	v := url.Values{}
	v.Set("ipversion", ipversion)
	v.Set("datacenter", datacenter)
	v.Set("platform", platform)

	req, err := c.NewRequest("POST", "ip/listfree", strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}

	var listfree IPListFree

	resp, err := c.Do(req, &listfree)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		fmt.Println("Unexpected status")
	}

	return &listfree, nil
}
