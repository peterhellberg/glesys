package glesys

type ChangelogResponse struct {
	Status    Status          `json:"status"`
	Changelog []ChangelogItem `json:"changelog"`
	Debug     Debug           `json:"debug"`
}

type ChangelogItem struct {
	Version string `json:"version"`
	Date    string `json:"date"`
	Notes   []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"notes"`
}

///////////////////////////////////////////////////////////////////////////////
// changelog/api
///////////////////////////////////////////////////////////////////////////////

func (c *Client) ChangelogAPI() (*ChangelogResponse, error) {
	return c.getChangelog("api")
}

///////////////////////////////////////////////////////////////////////////////
// changelog/controlpanel
///////////////////////////////////////////////////////////////////////////////

func (c *Client) ChangelogControlPanel() (*ChangelogResponse, error) {
	return c.getChangelog("controlpanel")
}

func (c *Client) getChangelog(section string) (*ChangelogResponse, error) {
	req, err := c.GetRequest("changelog/" + section)
	if err != nil {
		return nil, err
	}

	var r struct {
		Response ChangelogResponse `json:"response"`
	}

	if _, err = c.Do(req, &r); err != nil {
		return nil, err
	}

	return &r.Response, nil
}
