package utils

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func Shutdown(cancel context.CancelFunc) {
	sCh := make(chan os.Signal, 1)
	signal.Notify(sCh, os.Interrupt, syscall.SIGTERM)

	<-sCh
	cancel()
}
