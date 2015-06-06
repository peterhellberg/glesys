package glesys

import "testing"

///////////////////////////////////////////////////////////////////////////////
// user/details
///////////////////////////////////////////////////////////////////////////////

func TestUserDetails(t *testing.T) {
	r, err := c.UserDetails()
	if err != nil {
		t.Errorf(`unexpected error %v`, err)
	}

	if got, want := r.Status.Code, 401; got != want {
		t.Errorf(`r.Status.Code = %#v, want %#v`, got, want)
	}

	if got, want := r.Status.Text, "This function requires you to be logged in as a user. User credentials from user/login"; got != want {
		t.Errorf(`r.Status.Text = %#v, want %#v`, got, want)
	}
}
