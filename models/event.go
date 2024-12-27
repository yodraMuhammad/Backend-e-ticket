package models

import (
	"context"
	"time"
)

type Event struct {
	ID        string
	Name      string
	Location  string
	Date      time.Time
	CreatedAt time.Time
	UpdateAt  time.Time
}

type EventRepository interface {
	GetMany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, eventId string) (*Event, error)
	CreaOne(ctx context.Context, event Event) (*Event, error)
}
