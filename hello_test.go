package main

import "testing"

func TestHello(t *testing.T) {
	emptyResult := hello("Paulo")

	if emptyResult != "Hello Paulo" {
		t.Errorf("failed with %v", emptyResult)
	}
}
