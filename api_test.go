package glesys

import "testing"

///////////////////////////////////////////////////////////////////////////////
// api/maintenance
///////////////////////////////////////////////////////////////////////////////

func TestAPIMaintenance(t *testing.T) {
	r, err := c.APIMaintenance()
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := r.Status.Code, 200; got != want {
		t.Errorf(`r.Status.Code = %#v, want %#v`, got, want)
	}
}

///////////////////////////////////////////////////////////////////////////////
// api/serviceinfo
///////////////////////////////////////////////////////////////////////////////

func TestAPIServiceInfo(t *testing.T) {
	r, err := c.APIServiceInfo()
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := r.Status.Code, 200; got != want {
		t.Errorf(`r.Status.Code = %#v, want %#v`, got, want)
	}
}

///////////////////////////////////////////////////////////////////////////////
// api/listfunctions
///////////////////////////////////////////////////////////////////////////////

func TestAPIListFunctions(t *testing.T) {
	r, err := c.APIListFunctions()
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := r.Status.Code, 200; got != want {
		t.Errorf(`r.Status.Code = %#v, want %#v`, got, want)
	}

	if got, want := r.Modules.Account[1].Function, "info"; got != want {
		t.Errorf(`r.Modules.Account[1].Function = %#v, want %#v`, got, want)
	}
}
