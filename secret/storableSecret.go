package secret

import (
	"errors"

	"github.com/dchest/uniuri"
)

// StorableSecret this is something that can be stored and retrieved either
// until it has been viewed the maximum number of times, or if it's after the
// expiration date
type StorableSecret struct {
	ID             string
	Value          string
	RemainingViews int
}

func StoreNewSecret(s string, numViews int) (StorableSecret, error) {
	newSecret := StorableSecret{
		ID: uniuri.NewLen(32),
	}

	if s == "" {
		return newSecret, errors.New("You can not create an empty secret")
	}

	newSecret.Value = s
	newSecret.RemainingViews = numViews
	return newSecret, nil
}
