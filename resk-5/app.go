package resk

import (
	_ "resk-5/apis/web"
	_ "resk-5/core/accounts"
	"resk-5/infra"
	"resk-5/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
}