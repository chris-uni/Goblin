package lexer

import (
	"fmt"
	"testing"
)

func Test_BinaryOperators(t *testing.T) {

	out, _, err := Lex("+-*/")
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[{BinaryOperator + 1 0} {BinaryOperator - 1 1} {BinaryOperator * 1 2} {BinaryOperator / 1 3} {EOF EOF 1 5}]"
	tokens := fmt.Sprintf("%v", out)

	if tokens != want {

		t.Errorf("got %v\nwant %v", tokens, want)
	}
}

func Test_ContitionalOperators(t *testing.T) {

	out, _, err := Lex("==!=<=>=++--<>")
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[{== == 1 0} {!= != 1 2} {ConditionalOperator <= 1 4} {ConditionalOperator >= 1 6} {ConditionalOperator ++ 1 8} {ConditionalOperator -- 1 10} {ConditionalOperator < 1 12} {ConditionalOperator > 1 13} {EOF EOF 1 15}]"
	tokens := fmt.Sprintf("%v", out)

	if tokens != want {

		t.Errorf("got %v\nwant %v", tokens, want)
	}
}

func Test_ShorthandOperators(t *testing.T) {

	out, _, err := Lex("+=-=*=/=")
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[{ShorthandOperator += 1 0} {ShorthandOperator -= 1 2} {ShorthandOperator *= 1 4} {ShorthandOperator /= 1 6} {EOF EOF 1 9}]"
	tokens := fmt.Sprintf("%v", out)

	if tokens != want {

		t.Errorf("got %v\nwant %v", tokens, want)
	}
}

func Test_Symbols(t *testing.T) {

	out, _, err := Lex("?(){}[].:=,;")
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[{? ? 1 0} {( ( 1 1} {) ) 1 2} {{ { 1 3} {} } 1 4} {[ [ 1 5} {] ] 1 6} {. . 1 7} {: : 1 8} {= = 1 9} {, , 1 10} {; ; 1 12} {EOF EOF 1 13}]"
	tokens := fmt.Sprintf("%v", out)

	if tokens != want {

		t.Errorf("got %v\nwant %v", tokens, want)
	}
}

func Test_BooleanKeywords(t *testing.T) {

	out, _, err := Lex("true false")
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[{Boolean true 1 5} {Boolean false 1 11} {EOF EOF 1 11}]"
	tokens := fmt.Sprintf("%v", out)

	if tokens != want {

		t.Errorf("got %v\nwant %v", tokens, want)
	}
}

func Test_LanguageKeywords(t *testing.T) {

	out, _, err := Lex("let const fn if else while for using")
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[{Let let 1 4} {Const const 1 10} {Fn fn 1 13} {If if 1 16} {Else else 1 21} {While while 1 27} {For for 1 31} {Using using 1 37} {EOF EOF 1 37}]"
	tokens := fmt.Sprintf("%v", out)

	if tokens != want {

		t.Errorf("got %v\nwant %v", tokens, want)
	}
}

func Test_VariableAssignment(t *testing.T) {

	var tests = []struct {
		in   string
		want string
	}{
		{"let i = 0;", "[{Let let 1 4} {Identifier i 1 6} {= = 1 6} {Number 0 1 10} {; ; 1 10} {EOF EOF 1 11}]"},
		{"const j = 10;", "[{Const const 1 6} {Identifier j 1 8} {= = 1 8} {Number 10 1 13} {; ; 1 13} {EOF EOF 1 14}]"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.in)
		t.Run(testname, func(t *testing.T) {

			out, _, err := Lex(tt.in)
			tokens := fmt.Sprintf("%v", out)

			if err != nil {
				t.Errorf(err.Error())
			}

			if tokens != tt.want {
				t.Errorf("\ngot %v\nwant %v", tokens, tt.want)
			}
		})
	}
}
