package expeval

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{
			expr: "sqrt(A/pi)",
			env: Env{
				"A":  87616,
				"pi": math.Pi,
			},
			want: "167",
		},
		{
			expr: "pow(x,3)+pow(y,3)",
			env: Env{
				"x": 12,
				"y": 1,
			},
			want: "1729",
		},
		{
			expr: "pow(x,3)+pow(y,3)",
			env: Env{
				"x": 9,
				"y": 10,
			},
			want: "1729",
		},
		{
			expr: "5/9*(F-32)",
			env: Env{
				"F": -40,
			},
			want: "-40",
		},
		{
			expr: "5/9*(F-32)",
			env: Env{
				"F": 32,
			},
			want: "0",
		},
		{
			expr: "5/9*(F-32)",
			env: Env{
				"F": 212,
			},
			want: "100",
		},
	}

	var prevExpr string
	for _, test := range tests {
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n", test.expr, test.env, got, test.want)
		}
	}
}
