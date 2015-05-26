package glesys

import "testing"

///////////////////////////////////////////////////////////////////////////////
// server/list
///////////////////////////////////////////////////////////////////////////////

func TestServerList(t *testing.T) {
	r, err := c.ServerList()
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := r.Status.Code, 200; got != want {
		t.Errorf(`r.Status.Code = %#v, want %#v`, got, want)
	}
}
