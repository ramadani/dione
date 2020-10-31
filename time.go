package dione

import "time"

type goTime struct {
	expiresIn time.Duration
}

func (e *goTime) ExpiresAt() time.Time {
	return time.Now().Add(e.expiresIn)
}

func (e *goTime) ExpiresIn() time.Duration {
	return e.expiresIn
}

func NewTimeExpirer(expiresIn time.Duration) Expirer {
	return &goTime{expiresIn}
}
