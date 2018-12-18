package model

import (
	etcd_client "github.com/coreos/etcd/clientv3"
	"github.com/jmoiron/sqlx"
)

var (
	Db             *sqlx.DB
	EtcdClient     *etcd_client.Client
	EtcdPrefix     string
	EtcdProductKey string
)

func Init(db *sqlx.DB, etcdClient *etcd_client.Client, prefix, productKey string) (err error) {
	Db = db
	EtcdClient = etcdClient
	EtcdPrefix = prefix
	EtcdProductKey = productKey
	return
}
