package svc

import (
	"{{.Module}}/config"

	etcdv3 "go.etcd.io/etcd/client/v3"
)

type Svc struct {
	Config *config.Config
	Etcd   *etcdv3.Client
}

func newEtcdClient(config *config.Config) *etcdv3.Client {
	client, err := etcdv3.NewFromURLs(config.Etcd.Hosts)
	if err != nil {
		panic(err)
	}
	return client
}

func NewSvc(conf *config.Config) *Svc {
	client := newEtcdClient(conf)
	return &Svc{
		Config: conf,
		Etcd:   client,
	}
}