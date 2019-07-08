package secret

import "time"

// StorableSecret this is something that can be stored and retrieved either
// until it has been viewed the maximum number of times, or if it's after the
// expiration date
type StorableSecret struct {
	Value          string
	Expires        time.Time
	RemainingViews int
}
