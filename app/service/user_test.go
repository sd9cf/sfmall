package service

import (
	"testing"
)

func TestGetProfile(t *testing.T) {
	profile, err := User.GetProfile("oUT385ZLmRr6R_a9xKSfSW9SekYI")
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(profile)
}