package random

import (
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// String generates a random alphanumeric string
func String() string {
	return strconv.FormatInt(rand.Int63(), 36)
}
