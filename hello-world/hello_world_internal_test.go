package main

import "testing"

func TestHelloWorld(t *testing.T){

want := "Hello World!"
got := greet()

if got != want{

t.Errorf("Want: %v, got: %v", want, got)

}

}
