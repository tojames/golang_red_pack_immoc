package gorpc

import (
	"git.imooc.com/wendell1000/infra"
	"git.imooc.com/wendell1000/infra/base"
)

type GoRpcApiStarter struct {
	infra.BaseStarter
}

func (g *GoRpcApiStarter) Init(ctx infra.StarterContext) {
	base.RpcRegister(new(EnvelopeRpc))
}
