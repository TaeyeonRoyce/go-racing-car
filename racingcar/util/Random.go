package util

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber() int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(10)
}
