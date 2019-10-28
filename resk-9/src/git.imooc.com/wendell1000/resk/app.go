package resk

import (
	_ "git.imooc.com/wendell1000/account/core/accounts"
	"git.imooc.com/wendell1000/infra"
	"git.imooc.com/wendell1000/infra/base"
	"git.imooc.com/wendell1000/resk/apis/gorpc"
	_ "git.imooc.com/wendell1000/resk/apis/gorpc"
	_ "git.imooc.com/wendell1000/resk/apis/web"
	_ "git.imooc.com/wendell1000/resk/core/envelopes"
	"git.imooc.com/wendell1000/resk/jobs"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.GoRPCStarter{})
	infra.Register(&gorpc.GoRpcApiStarter{})
	infra.Register(&jobs.RefundExpiredJobStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
