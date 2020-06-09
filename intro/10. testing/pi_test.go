package main

import (
	"reflect"
	"testing"
)

func TestGetpi(t *testing.T) {
	result := Getpi()

	resultType := reflect.TypeOf(result).Kind()
	if resultType.String() != "float64" {
		t.Errorf("Type %T is not a float64", result)
	}
}

func TestGetpiformatted(t *testing.T) {
	result := Getpiformatter()

	if result != "3.1415926535897931" {
		t.Errorf("Result %v is not 3.1415926535897931", result)
	}
}
