package glesys

import "net/url"

///////////////////////////////////////////////////////////////////////////////
// customer/contactinfo
///////////////////////////////////////////////////////////////////////////////

type CustomerContactInfo struct {
	Response struct {
		Status      Status `json:"status"`
		ContactInfo struct {
			CustomerNumber string `json:"customernumber"`
			Contact        struct {
				CompanyName      string `json:"companyname"`
				ContactPerson    string `json:"contactperson"`
				InvoiceReference string `json:"InvoiceReference"`
				Address          string `json:"address"`
				ZipCode          string `json:"zipcode"`
				City             string `json:"city"`
				Country          string `json:"country"`
				NationalIDNumber string `json:"nationalidnumber"`
				PhoneNumber      string `json:"phonenumber"`
				Email            string `json:"email"`
			} `json:"contact"`
			Invoice struct {
				Address string `json:"invoiceaddress"`
				ZipCode string `json:"invoicezipcode"`
				City    string `json:"invoicecity"`
				Country string `json:"invoicecountry"`
			} `json:"invoice"`
			Settings struct {
				SeparateInvoiceAddress string `json:"separateinvoiceaddress"`
			} `json:"settings"`
		} `json:"contactinfo"`
		Debug Debug `json:"debug"`
	} `json:"response"`
}

func (c *Client) CustomerContactInfo(args ...func(*url.Values)) (*CustomerContactInfo, error) {
	data := &url.Values{}

	for _, f := range args {
		f(data)
	}

	req, err := c.PostRequest("customer/contactinfo", *data)
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
