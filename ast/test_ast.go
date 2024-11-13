package ast

import (
	"testing"

	"github.com/Rayan-BA/monkey/token"
)

func TestString(t *testing.T) {
	// test 'let x = y;'
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "x"},
					Value: "x",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "y"},
					Value: "y",
				},
			},
		},
	}

	if program.String() != "let x = y;" {
		t.Errorf("program.String() error. got=%q", program.String())
	}
}
