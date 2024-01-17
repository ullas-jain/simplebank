package token

import (
	"time"
)

// Maker ...
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(tokenString string) (*Payload, error)
}
