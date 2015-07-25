package io

import(
	"testing"
)

func TestStdio(t *testing.T){
	
	stdin := NewStdIn()

	if(!stdin.IsEmpty()){
		t.Errorf("boom")
	}
}

