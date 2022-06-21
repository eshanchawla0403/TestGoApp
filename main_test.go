package main

import (
	"testing"
)

func TestReturnWorld(t *testing.T) {
	output := ReturnWorld()
	if output != "World!!!" {
		t.Error("The output isn't correct")
	}
}
