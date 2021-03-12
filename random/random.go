package random

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// String generates a random alphanumeric string
func String() string {
	return strconv.FormatInt(rand.Int63(), 36)
}

// LongString generates a long random alphanumeric string
func LongString() string {
	var sb strings.Builder
	for i := 0; i < 5; i++ {
		sb.WriteString(String())
	}
	return sb.String()
}
