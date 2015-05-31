package glesys

///////////////////////////////////////////////////////////////////////////////
// contactperson/list
///////////////////////////////////////////////////////////////////////////////

type ContactPersonListResponse struct {
	Status         Status          `json:"status"`
	ContactPersons []ContactPerson `json:"contactpersons"`
	Debug          Debug           `json:"debug"`
}

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
