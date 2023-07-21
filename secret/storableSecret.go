package secret

import (
	"errors"

	"github.com/dchest/uniuri"

	"time"
)

// StorableSecret this is something that can be stored and retrieved either
// until it has been viewed the maximum number of times, or if it's after the
// expiration date
type StorableSecret struct {
	ID             string
	Value          string
	RemainingViews int
	EndTime        time.Time
	TimeString     string
}

func StoreNewSecret(s string, numViews int, addTime int) (StorableSecret, error) {
	newSecret := StorableSecret{
		ID: uniuri.NewLen(32),
	}

	if s == "" {
		return newSecret, errors.New("You can not create an empty secret")
	}

	newSecret.Value = s
	newSecret.RemainingViews = numViews
	newSecret.EndTime = time.Now().Add(time.Hour * time.Duration(addTime))
	newSecret.TimeString = newSecret.EndTime.Format(time.RFC822)
	return newSecret, nil
}

func getTime(askdf time.Time) (time.Time, error) {
	return time.Now(), nil
}
