package libs

import (
	"math/rand"
	"time"
)

type RandLib struct {
	Charset string
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func (rl RandLib) StringWithCharset(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = rl.Charset[seededRand.Intn(len(rl.Charset))]
	}
	return string(b)
}
