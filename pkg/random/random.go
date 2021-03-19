package random

import (
	"math/rand"
	"time"
)

var (
	l   = []rune("abcdefghijklmnopqrstuvwxyz")
	lun = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// String generates a random lowercase string with given length
func LowercaseString(length int) string {
	return randomString(length, l)
}

// String generates a random alphanumeric string with given length
func String(length int) string {
	return randomString(length, lun)
}

func randomString(length int, src []rune) string {
	r := make([]rune, length)
	for i := range r {
		r[i] = src[rand.Intn(len(src))]
	}
	return string(r)
}
