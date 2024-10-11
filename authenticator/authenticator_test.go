package authenticator

import (
	"testing"
)

func TestIsLoginCorrect(t *testing.T) {
	username := "EldonMacDonald"
	password := "this*isPassword1"

	auth := New()

	success, err := auth.IsLoginCorrect(username, password)
	if !success || err != nil {
		t.Fatal(err.Error())
	}
}
