package main

import "testing"

func TestShouldTrueIfUserIsAnAdult(t *testing.T) {
	user, _ := NewUser("Spider-man", 30, "spiderman@marvel.com")
	result := user.isAdult()
	assert(t, result, true)
}

func TestShouldFalseIfUserIsNotAnAdult(t *testing.T) {
	user, _ := NewUser("jack-jack", 3, "jack-jack@disney.com")
	result := user.isAdult()
	assert(t, result, false)
}

func assert(t *testing.T, result bool, expectedResult bool) {
	if result != expectedResult {
		t.Log("Result", result)
		t.Fail()
	}
}
