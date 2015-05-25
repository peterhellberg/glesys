package glesys

type ChangelogAPI struct {
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

func (c *Client) ChangelogAPI() (*ChangelogAPI, error) {
	req, err := c.GetRequest("changelog/api")
	if err != nil {
		return nil, err
	}

	var api ChangelogAPI

	if _, err = c.Do(req, &api); err != nil {
		return nil, err
	}

	return &api, nil
}
