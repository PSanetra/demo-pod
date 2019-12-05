package readiness

import "time"

type Readiness struct {
	ReadyAfter time.Time `json:"ready_after"`
}
