package main

import (
	"flag"
	"git.imooc.com/wendell1000/infra"
	"git.imooc.com/wendell1000/infra/base"
	_ "git.imooc.com/wendell1000/resk"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/tietang/props/consul"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	//通过HTTP服务来开启运行时性能剖析
	go func() {
		log.Info(http.ListenAndServe(":6060", nil))
	}()

	flag.Parse()
	profile := flag.Arg(0)
	if profile == "" {
		profile = "dev"
	}
	//获取程序运行文件所在的路径
	file := kvs.GetCurrentFilePath("boot.ini", 1)
	log.Info(file)
	//加载和解析配置文件
	conf := ini.NewIniFileCompositeConfigSource(file)
	if _, e := conf.Get("profile"); e != nil {
		conf.Set("profile", profile)
	}

	addr := conf.GetDefault("consul.address", "127.0.0.1:8500")
	contexts := conf.KeyValue("consul.contexts").Strings()
	log.Info("consul address:", addr)
	log.Info("consul contexts:", contexts)
	consulConf := consul.NewCompositeConsulConfigSourceByType(contexts, addr, kvs.ContentIni)
	consulConf.Add(conf)

	//
	base.InitLog(consulConf)
	app := infra.New(consulConf)
	app.Start()
}
