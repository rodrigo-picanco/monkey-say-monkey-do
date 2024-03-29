package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
                                Value: &Identifier{
                                  Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
                                  Value: "anotherVar",
                                },
			},
		},
	}

        if program.String() != "let myVar = anotherVar;" {
          t.Errorf("program.String() wrong. got=%q", program.String())
        }
}

func TestBoolean(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
                                Value: &Identifier{
                                  Token: token.Token{Type: token.TRUE, Literal: "true"},
                                  Value: "true",
                                },
			},
		},
	}

        if program.String() != "let myVar = true;" {
          t.Errorf("program.String() wrong. got=%q", program.String())
        }

}
