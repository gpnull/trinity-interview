package errors

import (
	"errors"
)

var (
	ErrInvalidToken    = errors.New("invalid token")
	ErrInvalidCampaign = errors.New("campaign is expired or invalid")
)
