package main

import "testing"

func TestGreeting(t *testing.T) {
	result := greeting("David")
	expect := "Hi, David"

	if result != expect {
		t.Errorf("result '%s', expect '%s'", result, expect)
	}

}