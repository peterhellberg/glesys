package glesys

import "testing"

///////////////////////////////////////////////////////////////////////////////
// contactperson/list
///////////////////////////////////////////////////////////////////////////////

func TestContactPersonList(t *testing.T) {
	r, err := c.ContactPersonList()
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := r.Status.Code, 200; got != want {
		t.Errorf(`r.Status.Code = %#v, want %#v`, got, want)
	}
}
