package generators

import (
	"math/rand"
)

func randInt(start, finish int) int {
	return rand.Intn(finish-start+1) + start
}

func Array(start, finish, qty int) []int {
	var arr []int
	for i := 0; i < qty; i++ {
		arr = append(arr, randInt(start, finish))
	}
	return arr
}
