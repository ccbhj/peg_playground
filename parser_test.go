package playground

import (
	"testing"
)

func TestPrintTree(t *testing.T) {
	if err := PrintAST(`10_2_3`); err != nil {
		t.Fatal(err)
	}

	if err := PrintAST(`0x123`); err != nil {
		t.Fatal(err)
	}
	println("-----------------------------------------------------")
}
