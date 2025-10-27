package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name" validate:"required,min=3"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	Password  string             `json:"password" bson:"password" validate:"required,min=6"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
	UpdatedAt int64              `json:"updated_at" bson:"updated_at"`
}

type UserResponse struct {
	ID        primitive.ObjectID `json:"_id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Email     string             `json:"email,omitempty"`
	CreatedAt int64              `json:"created_at,omitempty"`
} 