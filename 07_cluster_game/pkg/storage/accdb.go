package storage

import (
	"context"

	"github.com/topfreegames/pitaya/v2"

	"learn-pitaya-with-demos/cluster_game/pkg/models"
)

type AccountDB struct {
	app pitaya.Pitaya
}

// implement repo interface
var _ models.AccountRepo = (*AccountDB)(nil)

func NewAccountDB(app pitaya.Pitaya) *AccountDB {
	return &AccountDB{
		app: app,
	}
}

func (a AccountDB) GetByID(ctx context.Context, id string) (*models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountDB) GetByName(ctx context.Context, name string) (*models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountDB) Create(ctx context.Context, acc *models.Account) (*models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountDB) ChangePass(ctx context.Context, id, oldPass, newPass string) error {
	//TODO implement me
	panic("implement me")
}

func (a AccountDB) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
