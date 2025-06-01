package model

import (
	"context"
	"time"
)

type Message struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type MessageRepository interface {
	Create(ctx context.Context, message *Message) error
	GetRecent(ctx context.Context, limit int) ([]*Message, error)
	DeleteOld(ctx context.Context, maxAge time.Duration) error
}
