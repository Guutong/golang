package golang

import "testing"

func TestFactorial1(t *testing.T) {
	result := factorial(3)
	if result != 6 {
		t.Error("Please make the factorial function")
	}
}

func TestFactorial2(t *testing.T) {
	result := factorial(4)
	if result != 24 {
		t.Error("Please make the factorial function")
	}
}

func TestFactorial3(t *testing.T) {
	result := factorial(5)
	if result != 120 {
		t.Error("Please make the factorial function")
	}
}

func TestFactorial4(t *testing.T) {
	result := factorial(6)
	if result != 720 {
		t.Error("Please make the factorial function")
	}
}

func TestFactorial0(t *testing.T) {
	result := factorial(0)
	if result != 1 {
		t.Error("Please make the factorial function")
	}
}

func TestFactorialRecursive1(t *testing.T) {
	result := factorialRecursive(3)
	if result != 6 {
		t.Error("Please make the factorialRecursive function")
	}
}

func TestFactorialRecursive2(t *testing.T) {
	result := factorialRecursive(4)
	if result != 24 {
		t.Error("Please make the factorialRecursive function")
	}
}

func TestFactorialRecursive3(t *testing.T) {
	result := factorialRecursive(5)
	if result != 120 {
		t.Error("Please make the factorialRecursive function")
	}
}

func TestFactorialRecursive4(t *testing.T) {
	result := factorialRecursive(6)
	if result != 720 {
		t.Error("Please make the factorialRecursive function")
	}
}

func TestFactorialRecursive0(t *testing.T) {
	result := factorialRecursive(0)
	if result != 1 {
		t.Error("Please make the factorialRecursive function")
	}
}
