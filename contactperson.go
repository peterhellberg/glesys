package glesys

///////////////////////////////////////////////////////////////////////////////
// contactperson/list
///////////////////////////////////////////////////////////////////////////////

// ContactPersonListResponse contains the fields in a response from contactperson/list
type ContactPersonListResponse struct {
	Status         Status          `json:"status"`
	ContactPersons []ContactPerson `json:"contactpersons"`
	Debug          Debug           `json:"debug"`
}

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
