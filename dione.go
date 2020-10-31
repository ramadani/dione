package dione

import "time"

type Expirer interface {
	ExpiresAt() time.Time
	ExpiresIn() time.Duration
}
