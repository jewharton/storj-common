// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package sync2

import (
	"time"
)

// WithTimeout calls `do` concurrently and waits for it to complete. If the timeout
// is reached before `do` returns, `onTimeout` will be called; otherwise, `onTimeout`
// will not be called.
//
// Avoid attempting to detect whether a timeout has occurred from within `do`.
// Because `do` runs concurrently with the timeout timer, it may complete at the
// same time as the timeout timer expires, making detection within `do` unreliable.
// Logic specific to the timeout should instead be placed in `onTimeout`.
func WithTimeout(timeout time.Duration, do, onTimeout func()) {
	done := make(chan struct{})
	go func() {
		do()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(timeout):
		onTimeout()
		<-done
	}
}
