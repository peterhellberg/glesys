package glesys

import "net/url"

// ContactPerson is the fields for a contact person
type ContactPerson struct {
	ID                    int    `json:"contactpersonid"`
	Name                  string `json:"name"`
	Email                 string `json:"email"`
	PhoneNumber           string `json:"phonenumber"`
	PIN                   string `json:"personalidentificationnumber"`
	Role                  string `json:"role"`
	Account               string `json:"account"`
	ReceiveServiceNotices string `json:"receiveservicenotices"`
}

// ContactPersonResponse containst the fields in a response for a single contactperson
type ContactPersonResponse struct {
	Status        Status        `json:"status"`
	ContactPerson ContactPerson `json:"contactperson"`
	Debug         Debug         `json:"debug"`
}

///////////////////////////////////////////////////////////////////////////////
// contactperson/list
///////////////////////////////////////////////////////////////////////////////

// ContactPersonListResponse contains the fields in a response from contactperson/list
type ContactPersonListResponse struct {
	Status         Status          `json:"status"`
	ContactPersons []ContactPerson `json:"contactpersons"`
	Debug          Debug           `json:"debug"`
}

// ContactPersonList lists all the contact persons for this customer.
func (c *Client) ContactPersonList() (*ContactPersonListResponse, error) {
	req, err := c.GetRequest("contactperson/list")
	if err != nil {
		return nil, err
	}

	var r struct {
		Response ContactPersonListResponse `json:"response"`
	}

	if _, err = c.Do(req, &r); err != nil {
		return nil, err
	}

	return &r.Response, nil
}

///////////////////////////////////////////////////////////////////////////////
// contactperson/add
///////////////////////////////////////////////////////////////////////////////

func (c *Client) ContactPersonAdd(args ...func(*url.Values)) (*ContactPersonResponse, error) {
	data := &url.Values{}

	for _, f := range args {
		f(data)
	}

	req, err := c.PostRequest("contactperson/add", *data)
	if err != nil {
		return nil, err
	}

	var r struct {
		Response ContactPersonResponse `json:"response"`
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

///////////////////////////////////////////////////////////////////////////////
// contactperson/edit
///////////////////////////////////////////////////////////////////////////////

func (c *Client) ContactPersonEdit(id string, args ...func(*url.Values)) (*ContactPersonResponse, error) {
	data := &url.Values{}

	for _, f := range args {
		f(data)
	}

	data.Set("contactpersonid", id)

	req, err := c.PostRequest("contactperson/edit", *data)
	if err != nil {
		return nil, err
	}

	var r struct {
		Response ContactPersonResponse `json:"response"`
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

///////////////////////////////////////////////////////////////////////////////
// contactperson/delete
///////////////////////////////////////////////////////////////////////////////

func (c *Client) ContactPersonDelete(id string) (*ContactPersonResponse, error) {
	req, err := c.PostRequest("contactperson/delete", url.Values{
		"contactpersonid": {id},
	})
	if err != nil {
		return nil, err
	}

	var r struct {
		Response ContactPersonResponse `json:"response"`
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
