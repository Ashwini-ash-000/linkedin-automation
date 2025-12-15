package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

var letters = "abcdefghijklmnopqrstuvwxyz"

// TypeLikeHuman types text with human-like rhythm and occasional typos
func TypeLikeHuman(el *rod.Element, text string) {
	for _, ch := range text {
		// 5% chance of typo
		if rand.Float64() < 0.05 {
			typo := string(letters[rand.Intn(len(letters))])
			el.MustInput(typo)
			Think(80*time.Millisecond, 150*time.Millisecond)
			el.MustInput("\b") // backspace
			Think(50*time.Millisecond, 120*time.Millisecond)
		}

		el.MustInput(string(ch))
		Think(60*time.Millisecond, 180*time.Millisecond)
	}
}
