package liveness

import "sync/atomic"

type Settings struct {
	alive atomic.Value
}

func NewSettings() *Settings {
	settings := Settings{}

	settings.alive.Store(true)

	return &settings
}
