package email

import (
	"testing"
)

func Test_ValidateAddress(t *testing.T) {
	// Valid addresses
	_, err := ValidateAddress("name@example.com")
	if err != nil {
		t.Error(err)
	}
	_, err = ValidateAddress("name@x.y.z.example.co.uk")
	if err != nil {
		t.Error(err)
	}
	_, err = ValidateAddress("name@example.aero")
	if err != nil {
		t.Error(err)
	}
	_, err = ValidateAddress("x@example.com")
	if err != nil {
		t.Error(err)
	}
	_, err = ValidateAddress("first.last@example.com")
	if err != nil {
		t.Error(err)
	}
	_, err = ValidateAddress("first.middle.last@example.com")
	if err != nil {
		t.Error(err)
	}
	_, err = ValidateAddress("x+y@example.com")
	if err != nil {
		t.Error(err)
	}
	_, err = ValidateAddress("Name@example.com")
	if err != nil {
		t.Error(err)
	}
	_, err = ValidateAddress("NAme@x.y.z.example.co.uk")
	if err != nil {
		t.Error(err)
	}
	_, err = ValidateAddress("NAME@example.aero")
	if err != nil {
		t.Error(err)
	}
	_, err = ValidateAddress("X@example.com")
	if err != nil {
		t.Error(err)
	}
	_, err = ValidateAddress("First.Last@example.com")
	if err != nil {
		t.Error(err)
	}
	_, err = ValidateAddress("FIRST.middle.LAST@example.com")
	if err != nil {
		t.Error(err)
	}
	_, err = ValidateAddress("X+Y@example.com")
	if err != nil {
		t.Error(err)
	}

	// Invalid adresses
	_, err = ValidateAddress("name@.com")
	if err == nil {
		t.Error("Invalid email address not recognized")
	}
	_, err = ValidateAddress("name@example.x")
	if err == nil {
		t.Error("Invalid email address not recognized")
	}
	_, err = ValidateAddress("name@example.")
	if err == nil {
		t.Error("Invalid email address not recognized")
	}
	_, err = ValidateAddress("@example.com")
	if err == nil {
		t.Error("Invalid email address not recognized")
	}
	_, err = ValidateAddress(".@example.com")
	if err == nil {
		t.Error("Invalid email address not recognized")
	}
}
