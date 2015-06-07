package glesys

import (
	"fmt"
	"net/url"
	"testing"
)

///////////////////////////////////////////////////////////////////////////////
// contactperson/list
///////////////////////////////////////////////////////////////////////////////

func TestContactPersonList(t *testing.T) {
	r, err := c.ContactPersonList()
	if err != nil {
		t.Fatalf(`unexpected error %v`, err)
	}

	if got, want := r.Status.Code, 200; got != want {
		t.Errorf(`r.Status.Code = %#v, want %#v`, got, want)
	}
}

///////////////////////////////////////////////////////////////////////////////
// contactperson/edit
///////////////////////////////////////////////////////////////////////////////

func TestContactPersonEdit(t *testing.T) {
	var (
		contactPersonID = "17008"
	)

	r, err := c.ContactPersonEdit(func(v *url.Values) {
		v.Set("contactpersonid", contactPersonID)
		v.Set("name", "From test!")
	})
	if err != nil {
		t.Fatalf(`unexpected error %v`, err)
	}

	if got, want := r.Status.Code, 200; got != want {
		t.Errorf(`r.Status.Code = %#v, want %#v`, got, want)
	}
}

///////////////////////////////////////////////////////////////////////////////
// contactperson/delete
///////////////////////////////////////////////////////////////////////////////

func TestContactPersonDelete(t *testing.T) {
	r, err := c.ContactPersonAdd(func(v *url.Values) {
		v.Set("name", "Test person")
		v.Set("email", "test@example.com")
		v.Set("role", "operations")
		v.Set("receiveservicenotices", "false")
	})
	if err != nil {
		t.Fatalf(`unexpected error %v`, err)
	}

	r, err = c.ContactPersonDelete(fmt.Sprintf("%d", r.ContactPerson.ID))
	if err != nil {
		t.Fatalf(`unexpected error %v`, err)
	}

	if got, want := r.Status.Code, 200; got != want {
		t.Errorf(`r.Status.Code = %#v, want %#v`, got, want)
	}
}
