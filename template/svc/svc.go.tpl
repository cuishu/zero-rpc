package svc

import (
	"github.com/go-redis/redis/v8"
	"{{.Module.Name}}/config"

	etcdv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Svc struct {
	Config *config.Config
	Etcd   *etcdv3.Client
	Logger *zap.SugaredLogger
	Redis  *redis.Client
	DB     *gorm.DB
}

func newEtcdClient(config *config.Config) *etcdv3.Client {
	client, err := etcdv3.NewFromURLs(config.Etcd.Hosts)
	if err != nil {
		panic(err)
	}
	return client
}

func newLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return logger.Sugar()
}

func NewSvc(conf *config.Config) *Svc {
	client := newEtcdClient(conf)
	return &Svc{
		Config: conf,
		Etcd:   client,
		Logger: newLogger(),
	}
}