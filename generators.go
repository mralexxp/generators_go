package generators

import (
	"math/rand"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

// Generates a random string from base62
func ShortURL(length int) (string, error) {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		num := rand.Intn(len(base62Chars))
		result[i] = base62Chars[num]
	}
	return string(result), nil

}
