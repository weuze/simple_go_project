package main

import "testing"

func TestSayTrue(t *testing.T){
	v := true
	if v == false{
		t.Error("The value should be true")
	}
}
