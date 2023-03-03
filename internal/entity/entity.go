package entity

import (
	"time"
)

type Music struct {
	Name     string
	Duration time.Duration
}

func NewMusic() *Music {
	return &Music{}
}
