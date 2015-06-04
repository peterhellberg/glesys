package glesys

///////////////////////////////////////////////////////////////////////////////
// api/maintenance
///////////////////////////////////////////////////////////////////////////////

// APIMaintenanceResponse contains the fields in a response from api/maintenance
type APIMaintenanceResponse struct {
	Status      Status        `json:"status"`
	Maintenance []interface{} `json:"maintenance"`
	Debug       Debug         `json:"debug"`
}

// APIMaintenance returns a list of GleSYS services that are in maintenance mode.
func (c *Client) APIMaintenance() (*APIMaintenanceResponse, error) {
	req, err := c.GetRequest("api/maintenance")
	if err != nil {
		return nil, err
	}

	var r struct {
		Response APIMaintenanceResponse `json:"response"`
	}

	if _, err = c.Do(req, &r); err != nil {
		return nil, err
	}

	return &r.Response, nil
}

///////////////////////////////////////////////////////////////////////////////
// api/serviceinfo
///////////////////////////////////////////////////////////////////////////////

// APIServiceInfoResponse contains the fields in a response from api/serviceinfo
type APIServiceInfoResponse struct {
	Status      Status `json:"status"`
	ServiceInfo struct {
		Current []serviceInfoItem `json:"current"`
		Past    []serviceInfoItem `json:"past"`
	} `json:"serviceinfo"`
	Debug Debug `json:"debug"`
}

type serviceInfoItem struct {
	ServiceStatusID int    `json:"servicestatusid"`
	DateStart       string `json:"date_start"`
	DateEnd         string `json:"date_end"`
	Name            string `json:"name"`
	Text            string `json:"text"`
	Updates         []struct {
		Time    string `json:"time"`
		Message string `json:"message"`
	} `json:"updates"`
}

// APIServiceInfo returns service information. This is also available on http://drift.glesys.se
func (c *Client) APIServiceInfo() (*APIServiceInfoResponse, error) {
	req, err := c.GetRequest("api/serviceinfo")
	if err != nil {
		return nil, err
	}

	var r struct {
		Response APIServiceInfoResponse `json:"response"`
	}

	if _, err = c.Do(req, &r); err != nil {
		return nil, err
	}

	return &r.Response, nil
}

///////////////////////////////////////////////////////////////////////////////
// api/listfunctions
///////////////////////////////////////////////////////////////////////////////

// APIListFunctionsResponse contains the fields in a response from api/listfunctions
type APIListFunctionsResponse struct {
	Status  Status `json:"status"`
	Modules struct {
		Account       []functionDoc `json:"account"`
		API           []functionDoc `json:"api"`
		Archive       []functionDoc `json:"archive"`
		Changelog     []functionDoc `json:"changelog"`
		ContactPerson []functionDoc `json:"contactperson"`
		Customer      []functionDoc `json:"customer"`
		Domain        []functionDoc `json:"domain"`
		Email         []functionDoc `json:"email"`
		Invoice       []functionDoc `json:"invoice"`
		IP            []functionDoc `json:"ip"`
		LiveChat      []functionDoc `json:"livechat"`
		Loadbalancer  []functionDoc `json:"loadbalancer"`
		PaymentCard   []functionDoc `json:"paymentcard"`
		Server        []functionDoc `json:"server"`
		SSHKey        []functionDoc `json:"sshkey"`
		Transaction   []functionDoc `json:"transaction"`
		User          []functionDoc `json:"user"`
		VPN           []functionDoc `json:"vpn"`
	} `json:"modules"`
	Debug Debug `json:"debug"`
}

type functionDoc struct {
	Function      string `json:"function"`
	Documentation string `json:"documentation"`
}

// APIListFunctions lists all functions in the GleSYS API.
func (c *Client) APIListFunctions() (*APIListFunctionsResponse, error) {
	req, err := c.GetRequest("api/listfunctions")
	if err != nil {
		return nil, err
	}

	var r struct {
		Response APIListFunctionsResponse `json:"response"`
	}

	if _, err = c.Do(req, &r); err != nil {
		return nil, err
	}

	return &r.Response, nil
}
