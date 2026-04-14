package main

//如果测试用例很多，还可以用一个类似表格去表示

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"a1", 2, 3, 5},
		{"a2", 2, -3, -1},
		{"a3", 2, 0, 2},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Add(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got",
					c.A, c.B, c.Expected, ans)
			}
		})
	}
}
