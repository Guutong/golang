package golang

// Fibonacci numbers
// Fn= Fn-1 + Fn-2
// F0 = 0, F1 = 1
func fibonacci(a int) []int {
	f0, f1 := 0, 1
	ret := []int{}
	for i := 0; i < a; i++ {
		f0, f1 = f0+f1, f0
		ret = append(ret, f1)
	}
	return ret
}
