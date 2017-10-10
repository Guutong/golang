package golang

func factorial(a int) int {
	ret := 1
	for i := 1; i <= a; i++ {
		ret = ret * i
	}
	return ret
}

func factorialRecursive(a int) int {
	ret := 1
	if a >= 1 {
		ret = a * factorialRecursive(a-1)
	}
	return ret
}
