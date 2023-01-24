package usecase

import (
	"math/rand"
)

const (
	base    uint64 = 63
	charSet        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_"
	btm            = 100000000000000000
)

func toBase63(number uint64) string {
	result := ""

	for number > 0 {
		r := number % base
		number /= base
		result = string(charSet[r]) + result
	}
	return result
}

func GenerateShortURL() string {
	nmb := uint64(rand.Int63n(btm*10-btm+1) + btm)
	return toBase63(nmb)
}
