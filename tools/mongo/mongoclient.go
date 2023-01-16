package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type (
	Client interface {
		CreateDatabase(name string) error
	}
	ClientImpl struct {
		client *mongo.Client
		cfg    *ClientConfig
	}

	ClientConfig struct {
		Hosts          []string
		Port           string
		User           string
		Password       string
		Timeout        int
		ConnectTimeout int
		Database       string
	}
)

func (c ClientImpl) CreateDatabase(name string) error {
	//TODO implement me
	panic("implement me")
}

func NewCQLClient(cfg *ClientConfig) (Client, error) {
	var err error
	client := new(ClientImpl)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.ClientOptions{
		Hosts: cfg.Hosts,
		Auth: &options.Credential{
			Username: cfg.User,
			Password: cfg.Password,
		},
	}
	client.client, err = mongo.Connect(ctx, &clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}
