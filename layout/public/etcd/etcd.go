package etcd

import (
	"github.com/cro4k/gms/layout/public/global"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var c *clientv3.Client

func init() {
	client, err := clientv3.NewFromURLs(global.C().Etcd.Endpoints)
	if err != nil {
		logrus.Fatal(err)
	}
	c = client
}

func CLI() *clientv3.Client {
	return c
}

func TMP(f func(c *clientv3.Client) error) error {
	client, err := clientv3.NewFromURLs(global.C().Etcd.Endpoints)
	if err != nil {
		return err
	}
	defer client.Close()
	return f(client)
}
