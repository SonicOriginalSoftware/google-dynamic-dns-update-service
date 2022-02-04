package lib

import "time"

// RegisterTicker creates a periodic ticker
func RegisterTicker(frequencyTime int) (ticker *time.Ticker) {
	ticker = time.NewTicker(time.Duration(frequencyTime) * time.Second)
	defer ticker.Stop()
	return ticker
}
