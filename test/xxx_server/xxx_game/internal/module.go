package internal

import (
	"github.com/wuyutaott/leaf/module"
	"github.com/wuyutaott/leaf/test/xxx_server/tools"
)

var (
	skeleton = tools.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}
