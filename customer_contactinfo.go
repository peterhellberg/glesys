package glesys

import "net/url"

type CustomerContactInfo struct {
	Response struct {
		Status      ResponseStatus `json:"status"`
		ContactInfo struct {
			CustomerNumber string `json:"customernumber"`
			Contact        struct {
				CompanyName string `json:"companyname"`
			} `json:"contact"`
		} `json:"contactinfo"`
		Debug ResponseDebug `json:"debug"`
	} `json:"response"`
}

func (c *Client) CustomerContactInfo(args ...url.Values) (*CustomerContactInfo, error) {
	req, err := c.PostRequest("customer/contactinfo", url.Values{})
	if err != nil {
		return nil, err
	}

	var contactInfo CustomerContactInfo

	resp, err := c.Do(req, &contactInfo)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return &contactInfo, ErrUnexpectedHTTPStatus
	}

	return &contactInfo, nil
}
