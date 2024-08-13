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

func TestValue(t *testing.T) {
	var cases = []struct {
		Desc string
		Expr string
	}{
		{
			Desc: "NIL",
			Expr: "nil",
		},
		{
			Desc: "boolean",
			Expr: "#t",
		},
		{
			Desc: "decimal integer",
			Expr: "123456",
		},
		{
			Desc: "hex integer",
			Expr: "0x1f1",
		},
		{
			Desc: "short string",
			Expr: `"hello\n\x1f\u1234\U12345678"`,
		},
		{
			Desc: "long string",
			Expr: `"""foobar"""`,
		},
		{
			Desc: "identifier",
			Expr: `"""hello\n\x1f\u1234\U12345678"""`,
		},
		{
			Desc: "expression",
			Expr: `(HELLO world 123 "!!!")`,
		},
	}

	for _, c := range cases {
		println("[TEST] " + c.Desc)

		if err := PrintAST(c.Expr); err != nil {
			t.Fatal(err)
		}
		println("-----------------------------------------------------")
	}
}
