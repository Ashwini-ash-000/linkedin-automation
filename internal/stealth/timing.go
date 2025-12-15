package stealth

import (
	"math/rand"
	"time"
)

// Think simulates human pause before an action
func Think(min, max time.Duration) {
	if max <= min {
		time.Sleep(min)
		return
	}

	jitter := rand.Int63n(int64(max-min)) + int64(min)
	time.Sleep(time.Duration(jitter))
}

// ShortPause simulates micro delays between small actions
func ShortPause() {
	Think(50*time.Millisecond, 150*time.Millisecond)
}

// LongPause simulates reading/thinking time
func LongPause() {
	Think(800*time.Millisecond, 2*time.Second)
}
