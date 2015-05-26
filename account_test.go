package glesys

import "testing"

///////////////////////////////////////////////////////////////////////////////
// account/info
///////////////////////////////////////////////////////////////////////////////

func TestAccountInfo(t *testing.T) {
	r, err := c.AccountInfo()
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := r.Status.Code, 200; got != want {
		t.Errorf(`r.Status.Code = %#v, want %#v`, got, want)
	}

	if got, want := r.AccountInfo.Currency, "EUR"; got != want {
		t.Errorf(`r.AccountInfo.Currency = %#v, want %#v`, got, want)
	}
}
