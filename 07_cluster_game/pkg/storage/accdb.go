package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"learn-pitaya-with-demos/cluster_game/pkg/models"
)

type AccountDB struct {
	*mongo.Client
}

// implement repo interface
var _ models.AccountRepo = (*AccountDB)(nil)

func NewAccountDB(client *mongo.Client) *AccountDB {
	return &AccountDB{
		Client: client,
	}
}

func (a *AccountDB) GetByID(ctx context.Context, id string) (*models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountDB) GetByName(ctx context.Context, name string) (*models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountDB) Create(ctx context.Context, acc *models.Account) (*models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountDB) ChangePass(ctx context.Context, id, oldPass, newPass string) error {
	//TODO implement me
	panic("implement me")
}

func (a *AccountDB) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
