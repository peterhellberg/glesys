package glesys

type ChangelogControlPanel struct {
	Response struct {
		Status    ResponseStatus `json:"status"`
		Changelog []struct {
			Version string `json:"version"`
			Date    string `json:"date"`
			Notes   []struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"notes"`
		} `json:"changelog"`
		Debug ResponseDebug `json:"debug"`
	} `json:"response"`
}

func (c *Client) ChangelogControlPanel() (*ChangelogControlPanel, error) {
	req, err := c.GetRequest("changelog/controlpanel")
	if err != nil {
		return nil, err
	}

	var controlPanel ChangelogControlPanel

	if _, err = c.Do(req, &controlPanel); err != nil {
		return nil, err
	}

	return &controlPanel, nil
}
