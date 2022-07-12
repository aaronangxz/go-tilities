package go_tilities

import (
	"math/rand"
	"sort"
	"time"
)

func Swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func MakeSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())
	return rand.Perm(size)
}

func MakeSortedSlice(size int) []int {
	s := MakeSlice(size)
	sort.Ints(s)
	return s
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ReverseSlice(s *[]int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
