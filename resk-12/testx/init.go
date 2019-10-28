package testx

import (
	"git.imooc.com/wendell1000/infra"
	"git.imooc.com/wendell1000/infra/base"
	"git.imooc.com/wendell1000/resk/apis/gorpc"
	"git.imooc.com/wendell1000/resk/core/accounts"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

func init() {

	//获取程序运行文件所在的路径
	file := kvs.GetCurrentFilePath("../brun/config.ini", 1)
	//加载和解析配置文件
	conf := ini.NewIniFileCompositeConfigSource(file)
	base.InitLog(conf)
	conf.Set("testing", "true")

	//

	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.GoRPCStarter{})
	infra.Register(&gorpc.GoRpcApiStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&accounts.AccountClientStarter{})
	infra.Register(&base.HookStarter{})

	app := infra.New(conf)
	app.Start()
	//time.Sleep(time.Second * 3)
}
