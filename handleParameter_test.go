package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestGetPort(t *testing.T){
    t.Run("ManualPort", func(t *testing.T) {
	expectedPort := "42069"
	var expectedErr error = nil

	got, err := GetPort("--PORT=42069")

	if !reflect.DeepEqual(expectedPort, got) && expectedErr != err {
	    t.Errorf("Expected port: '%s' -- Got: '%s'\nExpected nil err but got '%s'", expectedPort, got, err.Error())
	}
    })
    t.Run("LowerPort", func(t *testing.T){
	expectedPort := ""
	expectedErr := errors.New("Can't bind lower ports -- Received: 222")

	got, err := GetPort("--PORT=222")

	if !reflect.DeepEqual(expectedPort, got) && expectedErr.Error() != err.Error() {
	    t.Errorf("Expected port: '%s' -- Got: '%s'\nExpected err to be '%s' but got '%s'", expectedPort, got, expectedErr.Error(), err.Error())
	}
    })
    t.Run("HigherPort", func(t *testing.T){
	expectedPort := ""
	expectedErr := errors.New("Can't bind private ports -- Received: 69420")

	got, err := GetPort("--PORT=69420")

	if !reflect.DeepEqual(expectedPort, got) && expectedErr.Error() != err.Error() {
	    t.Errorf("Expected port: '%s' -- Got: '%s'\nExpected err: '%s' -- got: '%s'", expectedPort, got, expectedErr.Error(), err.Error())
	}
    })

    t.Run("InvalidPort", func(t *testing.T) {
	expectedPort := ""
	expectedErr := errors.New("The PORT must be a number ranging from 1,024 to 49,151, got: myPort")

	got, err := GetPort("--PORT=myPort")

	if !reflect.DeepEqual(expectedPort, got) && expectedErr.Error() != err.Error() {
	    t.Errorf("Expected port: '%s' -- Got: '%s'\nExpected err: '%s' -- got: '%s'", expectedPort, got, expectedErr.Error(), err.Error())
	}
    })

    t.Run("UnexpectedParam", func(t *testing.T) {
	expectedPort := ""
	expectedErr := errors.New("Unexpected parameter, expected --PORT=#### but received: --PART=2020")

	got, err := GetPort("--PART=2020")

	if !reflect.DeepEqual(expectedPort, got) && expectedErr.Error() != err.Error() {
	    t.Errorf("Expected port: '%s' -- Got: '%s'\nExpected err: '%s' -- got '%s'", expectedPort, got, expectedErr.Error(), err.Error())
	}
    })

    t.Run("MoreThanOneEqual", func(t *testing.T) {
	expectedPort := ""
	expectedErr := errors.New("Unexpected parameter, expected --PORT=#### but received: --PART=2020=")

	got, err := GetPort("--PORT=2020=")

	if !reflect.DeepEqual(expectedPort, got) && expectedErr.Error() != err.Error() {
	    t.Errorf("Expected port: '%s' -- Got: '%s'\nExpected err: '%s' -- got '%s'", expectedPort, got, expectedErr.Error(), err.Error())
	}
    })

    t.Run("EmptyPort", func(t *testing.T) {
	got, err := GetPort("--PoRt=")

	expectedPort:=""
	expectedErr:=errors.New("Empty PORT")

	if !reflect.DeepEqual(expectedPort, got) && expectedErr.Error() != err.Error() {
	    t.Errorf("Expected port: '%s' -- Got: '%s'\nExpected err: '%s' -- got '%s'", expectedPort, got, expectedErr.Error(), err.Error())
	}
    })
}



