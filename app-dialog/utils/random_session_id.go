package utils

import (
	"math/rand"
	"time"
)

func RandomSessionId() uint64 {
	rand.Seed(time.Now().Unix())
	now := uint64(time.Now().Unix())
	rnd := rand.Int31()
	return 0x8000000000000000 | (((now - 1577829600) << 31) + uint64(rnd))
}
