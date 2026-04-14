package main

//子测试:如果需要给一个函数，调用不同的测试用例，可以使用子测试
//子测试里面的Fatal，是不会终止程序的

import (
	"testing"
)

func TestAdd(t1 *testing.T) {
	t1.Run("add1", func(t *testing.T) {
		if ans := Add(1, 2); ans != 3 {
			// 如果不符合预期，那就是测试不通过
			t.Fatalf("1 + 2 expected be 3, but %d got", ans)
		}
	})
	t1.Run("add2", func(t *testing.T) {
		if ans := Add(-10, -20); ans != -30 {
			t.Fatalf("-10 + -20 expected be -30, but %d got", ans)
		}
	})

}

//如果测试用例很多，还可以用一个类似表格去表示 ——> 子测试2.go
