package models

import "context"

type Account struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Pass string `json:"pass" bson:"pass"`
}

type AccountRepo interface {
	GetByID(ctx context.Context, id string) (*Account, error)
	GetByName(ctx context.Context, name string) (*Account, error)
	Create(ctx context.Context, acc *Account) (*Account, error)
	ChangePass(ctx context.Context, id, oldPass, newPass string) error
	Delete(ctx context.Context, id string) error
}
