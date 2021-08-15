package handlers

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// TestGet get cache by key
// 11000001111010101000010010010000000000000000000000000000000000
// 00110000 01111010 10100001 01010101 1 0000000 00000000 00000000 00000000
// 9223372036854776000 -> 1000000000000000000000000000000000000000000000000000000000000000
// uint64 : 0 to 18446744073709551615
func TestHashAndSalt(t *testing.T) {
	var now = time.Now()
	snow := uint64(now.Unix())
	fmt.Println(snow)
	fmt.Println(snow << 31)
	rnd := rand.Int31()
	fmt.Println(rnd)
	fmt.Println(9223372036854776000 + snow<<31 + uint64(rnd))
}
