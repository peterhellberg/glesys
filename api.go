package glesys

///////////////////////////////////////////////////////////////////////////////
// api/maintenance
///////////////////////////////////////////////////////////////////////////////

type APIMaintenance struct {
	Status      Status        `json:"status"`
	Maintenance []interface{} `json:"maintenance"`
	Debug       Debug         `json:"debug"`
}

// APIMaintenance returns a list of GleSYS services that are in maintenance mode.
func (c *Client) APIMaintenance() (*APIMaintenance, error) {
	req, err := c.GetRequest("api/maintenance")
	if err != nil {
		return nil, err
	}

	var r struct {
		Response APIMaintenance `json:"response"`
	}

	if _, err = c.Do(req, &r); err != nil {
		return nil, err
	}

	return &r.Response, nil
}

///////////////////////////////////////////////////////////////////////////////
// api/serviceinfo
///////////////////////////////////////////////////////////////////////////////

type APIServiceInfo struct {
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

func (c *Client) APIServiceInfo() (*APIServiceInfo, error) {
	req, err := c.GetRequest("api/serviceinfo")
	if err != nil {
		return nil, err
	}

	var r struct {
		Response APIServiceInfo `json:"response"`
	}

	if _, err = c.Do(req, &r); err != nil {
		return nil, err
	}

	return &r.Response, nil
}

///////////////////////////////////////////////////////////////////////////////
// api/listfunctions
///////////////////////////////////////////////////////////////////////////////

type APIListFunctions struct {
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

func (c *Client) APIListFunctions() (*APIListFunctions, error) {
	req, err := c.GetRequest("api/listfunctions")
	if err != nil {
		return nil, err
	}

	var r struct {
		Response APIListFunctions `json:"response"`
	}

	if _, err = c.Do(req, &r); err != nil {
		return nil, err
	}

	return &r.Response, nil
}
