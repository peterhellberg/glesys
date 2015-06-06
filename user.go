package glesys

///////////////////////////////////////////////////////////////////////////////
// user/details
///////////////////////////////////////////////////////////////////////////////

// UserDetailsResponse contains the fields in a response from user/details
type UserDetailsResponse struct {
	Status Status `json:"status"`
	Debug  Debug  `json:"debug"`
}

// UserDetails provides information about the currently logged in user. Username, name, twofactor method and more.
func (c *Client) UserDetails() (*UserDetailsResponse, error) {
	req, err := c.GetRequest("user/details")
	if err != nil {
		return nil, err
	}

	var r struct {
		Response UserDetailsResponse `json:"response"`
	}

	if _, err = c.Do(req, &r); err != nil {
		return nil, err
	}

	return &r.Response, nil
}
