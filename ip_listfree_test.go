package glesys

import "testing"

func TestIPListFree(t *testing.T) {
	c := NewClient(username, apiKey, nil)

	var (
		ipversion  = "4"
		datacenter = "Stockholm"
		platform   = "OpenVZ"
	)

	r, err := c.IPListFree(ipversion, datacenter, platform)
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := r.Response.Status.Code, 200; got != want {
		t.Errorf(`r.Response.Status.Code = %#v, want %#v`, got, want)
	}

	if got, want := r.Response.IPList.Datacenter, datacenter; got != want {
		t.Errorf(`r.Response.IPList.Datacenter = %#v, want %#v`, got, want)
	}
}
