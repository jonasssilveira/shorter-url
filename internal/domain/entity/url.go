package entity

import (
	"time"
)

type URL struct {
	FullURL        string    `json:"full_url"`
	ExpirationDate time.Time `json:"expiration_date"`
}
