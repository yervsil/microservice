package repository

import (
	//	"github.com/yervsil/user_service/internal/user"

	"context"

	"github.com/yervsil/user_service/internal/user/domain"
	"github.com/yervsil/user_service/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	//"github.com/yervsil/user_service/internal/user/domain"
)

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db.Collection("user")}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) (string, error) {
	res, err := r.db.InsertOne(ctx, user)
	if mongodb.IsDuplicate(err) {
		return "", domain.ErrUserAlreadyExists
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (r *UserRepository) GetByCredentials(ctx context.Context, email, password string) (*domain.User, error) {
	var user *domain.User

	filter := bson.M{
		"email":         email,
		"password": 	password, // Предполагаем, что вы храните хэш пароля в поле "password_hash"
	}

	err := r.db.FindOne(ctx, filter).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Если документ с указанными данными не найден
			return nil, err
		}
		return nil, err
	}

	return user, nil
}

