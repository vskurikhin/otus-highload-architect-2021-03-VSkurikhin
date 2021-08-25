package utils

import (
	"math/rand"
	"time"
)

func RandomIdWithShardId(shardId uint8) uint64 {
	rand.Seed(time.Now().Unix())
	now := uint64(time.Now().UnixNano())
	rnd := (rand.Int31n(0x7FFFFF) << 8) + int32(shardId)
	return 0x8000000000000000 | (((now - 1577829600) << 31) + uint64(rnd))
}
