package glesys

import (
	"net/url"
	"testing"
)

func TestCustomerContactInfo(t *testing.T) {
	c := NewClient(username, apiKey, nil)

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

	if got, want := r.Response.Status.Code, 200; got != want {
		t.Errorf(`r.Response.Status.Code = %#v, want %#v`, got, want)
	}

	if got, want := r.Response.ContactInfo.CustomerNumber, customerNumber; got != want {
		t.Errorf(`r.Response.ContactInfo.CustomerNumber = %#v, want %#v`, got, want)
	}

	contact := r.Response.ContactInfo.Contact

	if got, want := contact.ZipCode, zipCode; got != want {
		t.Errorf(`r.Response.ContactInfo.CustomerNumber = %#v, want %#v`, got, want)
	}
}
