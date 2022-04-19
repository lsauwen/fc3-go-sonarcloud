package main

import "testing"

func TestSum(t *testing.T) {

	resul := sum(2, 3)

	if resul != 5 {
		t.Error("The result must be 5")
	}

}
