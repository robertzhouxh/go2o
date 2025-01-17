/**
 * Copyright 2015 @ 56x.net.
 * name : testing
 * author : jarryliu
 * date : 2016-06-15 08:31
 * description :
 * history :
 */
package ti

import (
	"github.com/ixre/go2o/core"
	"github.com/ixre/go2o/core/repos"
	"github.com/ixre/go2o/core/service/impl"
	"github.com/ixre/gof"
	"github.com/ixre/gof/db/orm"
	"go.etcd.io/etcd/clientv3"
	"os"
	"time"
)

var (
	Factory *repos.RepoFactory
)
var (
	REDIS_DB = "1"
)

func GetApp() gof.App {
	return gof.CurrentApp
}

func init() {
	// 默认的ETCD端点
	etcdEndPoints := []string{"http://127.0.0.1:2379"}
	cfg := clientv3.Config{
		Endpoints:   etcdEndPoints,
		DialTimeout: 5 * time.Second,
	}
	confPath := "app.conf"
	for {
		_, err := os.Stat(confPath)
		if err == nil {
			break
		}
		confPath = "../" + confPath
	}
	app := core.NewApp(confPath, &cfg)
	gof.CurrentApp = app
	core.Init(app, false, false)
	conn := app.Db()
	sto := app.Storage()
	o := orm.NewOrm(conn.Driver(), conn.Raw())
	impl.InitTestService(app, conn, o, sto)
	Factory = (&repos.RepoFactory{}).Init(o, sto)
}
