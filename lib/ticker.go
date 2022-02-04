package lib

import "time"

// RegisterTicker creates a periodic ticker
func RegisterTicker(frequencyTime int) (ticker *time.Ticker) {
	return time.NewTicker(time.Duration(frequencyTime) * time.Second)
}
