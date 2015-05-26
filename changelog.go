package glesys

type ChangelogResponse struct {
	Response struct {
		Status    Status          `json:"status"`
		Changelog []ChangelogItem `json:"changelog"`
		Debug     Debug           `json:"debug"`
	} `json:"response"`
}

type ChangelogItem struct {
	Version string `json:"version"`
	Date    string `json:"date"`
	Notes   []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"notes"`
}

func (c *Client) ChangelogAPI() (*ChangelogResponse, error) {
	return c.getChangelog("api")
}

func (c *Client) ChangelogControlPanel() (*ChangelogResponse, error) {
	return c.getChangelog("controlpanel")
}

func (c *Client) getChangelog(section string) (*ChangelogResponse, error) {
	req, err := c.GetRequest("changelog/" + section)
	if err != nil {
		return nil, err
	}

	var res ChangelogResponse

	if _, err = c.Do(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
