package golang

func concat(a, b []string) []string {
	for _, v := range b {
		a = append(a, v)
	}
	return a
}
