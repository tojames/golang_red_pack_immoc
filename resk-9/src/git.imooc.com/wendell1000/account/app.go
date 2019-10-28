package resk

import (
	_ "git.imooc.com/wendell1000/account/apis/web"
	_ "git.imooc.com/wendell1000/account/core/accounts"
	"git.imooc.com/wendell1000/infra"
	"git.imooc.com/wendell1000/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
