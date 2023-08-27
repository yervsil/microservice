package domain

import (
//	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	SessionID string    `json:"session_id"`
	UserID    primitive.ObjectID `json:"user_id"`
}
