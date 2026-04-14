package main

// 如果我们要实现一个对int类型的求和函数
func add(a, b int) int {
	return a + b
}

// 但是这样写了之后，如果参数是float类型，int32类型就没办法使用了
// 难道要为每个类型都写一个这样的函数吗？
// 显然这就不合理 ——> 泛型:
func add[T int | float64 | int32](a, b T) T {
	return a + b
}
