package resk

import (
	_ "imooc.com/resk/apis/web"
	_ "imooc.com/resk/core/accounts"
	"imooc.com/resk/infra"
	"imooc.com/resk/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
}
