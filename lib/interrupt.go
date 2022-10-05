package lib

import (
	"os"
	"os/signal"
	"syscall"
)

// RegisterInterruptHandler registers a channel for handling an program kill/interrupt signal
func RegisterInterruptHandler() (c chan os.Signal) {
	c = make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	return c
}
