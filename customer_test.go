package glesys

import (
	"net/url"
	"testing"
)

///////////////////////////////////////////////////////////////////////////////
// customer/contactinfo
///////////////////////////////////////////////////////////////////////////////

func TestCustomerContactInfo(t *testing.T) {
	var (
		customerNumber = "19454"
		zipCode        = "12432"
	)

	r, err := c.CustomerContactInfo(func(v *url.Values) {
		v.Set("zipcode", zipCode)
	})
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := r.Status.Code, 200; got != want {
		t.Errorf(`r.Status.Code = %#v, want %#v`, got, want)
	}

	if got, want := r.ContactInfo.CustomerNumber, customerNumber; got != want {
		t.Errorf(`r.ContactInfo.CustomerNumber = %#v, want %#v`, got, want)
	}

	contact := r.ContactInfo.Contact

	if got, want := contact.ZipCode, zipCode; got != want {
		t.Errorf(`contact.ZipCode = %#v, want %#v`, got, want)
	}
}
