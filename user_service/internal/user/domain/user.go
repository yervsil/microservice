package domain

import (
//	"github.com/yervsil/user_service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct {
	ID       primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password_hash,omitempty" bson:"password,omitempty"`
}