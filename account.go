package glesys

import "encoding/json"

///////////////////////////////////////////////////////////////////////////////
// account/info
///////////////////////////////////////////////////////////////////////////////

// AccountInfoResponse contains the fields in a response from account/info
type AccountInfoResponse struct {
	Status      Status `json:"status"`
	AccountInfo struct {
		AccountName    string `json:"accountname"`
		CustomerNumber int    `json:"customernumber"`
		Currency       string `json:"currency"`
		UnpaidInvoices string `json:"unpaidinvoices"`
		Locked         string `json:"locked"`
		LockedReason   string `json:"lockedreason"`
		APIEnabled     string `json:"apienabled"`
		ServiceLimits  struct {
			MaxNumServers              int    `json:"maxnumservers"`
			MaxNumIPv4                 int    `json:"maxnumipv4"`
			MaxNumIPv6                 int    `json:"maxnumipv6"`
			MaxArchiveTotalSize        int    `json:"maxarchivetotalsize"`
			MaxNumEmailAccounts        int    `json:"maxnumemailaccount"`
			CanChangeBillingPeriod     string `json:"canchangebillingperiod"`
			EmailGlobalQuota           int    `json:"emailglobalquota"`
			DomainRegistrationsAllowed int    `json:"domainregistrationsallowed"`
		} `json:"servicelimits"`
		Services struct {
			Server   string `json:"server"`
			IP       string `json:"ip"`
			Domain   string `json:"domain"`
			Archive  string `json:"archive"`
			Email    string `json:"email"`
			Invoice  string `json:"invoice"`
			Customer string `json:"customer"`
			Account  string `json:"account"`
			API      string `json:"api"`
			VPN      string `json:"vpn"`
		} `json:"services"`
		FeatureFlags []*json.RawMessage `json:"featureflags"`
	} `json:"accountinfo"`
	Debug Debug `json:"debug"`
}

// AccountInfo provides information about the currently logged in account.
func (c *Client) AccountInfo() (*AccountInfoResponse, error) {
	req, err := c.GetRequest("account/info")
	if err != nil {
		return nil, err
	}

	var r struct {
		Response AccountInfoResponse `json:"response"`
	}

	if _, err = c.Do(req, &r); err != nil {
		return nil, err
	}

	return &r.Response, nil
}
