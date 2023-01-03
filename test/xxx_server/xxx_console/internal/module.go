package internal

import (
	"github.com/wuyutaott/leaf/module"
	"github.com/wuyutaott/leaf/test/xxx_server/tools"
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = tools.NewSkeleton()
}

func (m *Module) OnDestroy() {

}
