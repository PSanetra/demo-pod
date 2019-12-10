package api

import (
	"demo-pod/api/liveness"
	"demo-pod/api/notes"
	"demo-pod/api/readiness"
	"demo-pod/api/watch"
	"time"
)

type Settings struct {
	StartupDelay      time.Duration
	CorsOrigins       []string
	LivenessSettings  *liveness.Settings
	NotesSettings     notes.Settings
	ReadinessSettings *readiness.Settings
	WatchSettings     watch.Settings
}
