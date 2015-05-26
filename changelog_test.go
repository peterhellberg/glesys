package glesys

import "testing"

func TestChangelogAPI(t *testing.T) {
	r, err := c.ChangelogAPI()
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := r.Response.Status.Code, 200; got != want {
		t.Errorf(`r.Response.Status.Code = %#v, want %#v`, got, want)
	}
}

func TestChangelogControlPanel(t *testing.T) {
	r, err := c.ChangelogControlPanel()
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := r.Response.Status.Code, 200; got != want {
		t.Errorf(`r.Response.Status.Code = %#v, want %#v`, got, want)
	}
}
