package readiness

import (
	"sync/atomic"
	"time"
)

type Settings struct {
	readyAfter atomic.Value
}

func NewSettings() *Settings {
	settings := Settings{}

	now := time.Now().UTC()

	settings.readyAfter.Store(&now)

	return &settings
}
