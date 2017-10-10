package golang

import (
	"strings"
)

func LettersCount(a string) map[string]int {
	ret := map[string]int{}
	strSplit := strings.Split(a, " ")
	for _, v := range strSplit {
		ret[v] = len(v)
	}
	return ret
}
