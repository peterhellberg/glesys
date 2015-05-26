package glesys

import "net/url"

///////////////////////////////////////////////////////////////////////////////
// customer/contactinfo
///////////////////////////////////////////////////////////////////////////////

type CustomerContactInfoResponse struct {
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
}

func (c *Client) CustomerContactInfo(args ...func(*url.Values)) (*CustomerContactInfoResponse, error) {
	data := &url.Values{}

	for _, f := range args {
		f(data)
	}

	req, err := c.PostRequest("customer/contactinfo", *data)
	if err != nil {
		return nil, err
	}

	var r struct {
		Response CustomerContactInfoResponse `json:"response"`
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
