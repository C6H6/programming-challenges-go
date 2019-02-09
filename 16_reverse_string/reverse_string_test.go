package main

import "testing"

func TestReverseReverses(t *testing.T) {
	if reverse("abcd") != "dcba" {
		t.Errorf("Invalid returned value abcd -> %s", reverse("abcd"))
	}
}
