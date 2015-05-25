package glesys

import "testing"

func TestChangelogControlPanel(t *testing.T) {
	c := NewClient(username, apiKey, nil)

	r, err := c.ChangelogControlPanel()
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := r.Response.Status.Code, 200; got != want {
		t.Errorf(`r.Response.Status.Code = %#v, want %#v`, got, want)
	}
}
