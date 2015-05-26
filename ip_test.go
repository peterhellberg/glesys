package glesys

import "testing"

///////////////////////////////////////////////////////////////////////////////
// ip/listfree
///////////////////////////////////////////////////////////////////////////////

func TestIPListFree(t *testing.T) {
	var (
		ipversion  = "4"
		datacenter = "Stockholm"
		platform   = "OpenVZ"
	)

	r, err := c.IPListFree(ipversion, datacenter, platform)
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := r.Status.Code, 200; got != want {
		t.Errorf(`r.Status.Code = %#v, want %#v`, got, want)
	}

	if got, want := r.IPList.Datacenter, datacenter; got != want {
		t.Errorf(`r.IPList.Datacenter = %#v, want %#v`, got, want)
	}
}
