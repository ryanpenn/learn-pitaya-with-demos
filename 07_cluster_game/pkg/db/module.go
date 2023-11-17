package db

import (
	"context"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoModule interface {
	GetClient() *mongo.Client
	GetName() string
}

type module struct {
	component.Base
	*mongo.Client
	name    string
	addr    string
	timeout time.Duration
}

func RegisterModule(app pitaya.Pitaya, name string, addr string, timeout time.Duration) error {
	m := &module{
		name:    name,
		addr:    addr,
		timeout: timeout,
	}
	return app.RegisterModule(m, name)
}

func GetModule(app pitaya.Pitaya, name string) MongoModule {
	if m, err := app.GetModule(name); err == nil {
		return m.(MongoModule)
	} else {
		logger.Log.WithError(err).Errorf("get module %s error", name)
		return nil
	}
}

func (m *module) GetClient() *mongo.Client {
	return m.Client
}

func (m *module) GetName() string {
	return m.name
}

func (m *module) Init() error {
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(m.addr).SetTimeout(m.timeout))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	if err = c.Ping(ctx, nil); err != nil {
		return err
	}

	m.Client = c
	return nil
}

func (m *module) Shutdown() error {
	if m.Client != nil {
		if err := m.Client.Disconnect(context.TODO()); err != nil {
			return err
		}
	}

	return nil
}
