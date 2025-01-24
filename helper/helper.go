package helper

import (
	"math/rand"
	"time"
)

// ShuffleString shuffles the characters in a string.
func ShuffleString(input string) string {
	rand.Seed(time.Now().UnixNano())
	runes := []rune(input)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}
