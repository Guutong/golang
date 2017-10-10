package golang

func concat(a, b []string) []string {
	// for _, v := range b {
	// 	a = append(a, v)
	// }
	// return a
	return append(a, b...)
}

func delFirst(a []string) []string {
	return a[1:]
}

func delLast(a []string) []string {
	return a[:len(a)-1]
}

func delSecond(a []string) []string {
	return append(a[:1], a[2:]...)
}

func odd(a []int) []int {
	ret := []int{}
	for _, v := range a {
		if v%2 != 0 {
			ret = append(ret, v)
		}
	}
	return ret
}
