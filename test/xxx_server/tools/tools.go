package tools

import (
	"github.com/wuyutaott/leaf/chanrpc"
	"github.com/wuyutaott/leaf/module"
	"github.com/wuyutaott/leaf/test/xxx_server/xxx_conf"
)

func NewSkeleton() *module.Skeleton {
	ske := &module.Skeleton{
		GoLen:              xxx_conf.GoLen,
		TimerDispatcherLen: xxx_conf.TimerDispatcherLen,
		AsynCallLen:        xxx_conf.AsynCallLen,
		ChanRPCServer:      chanrpc.NewServer(xxx_conf.ChanRPCLen),
	}
	ske.Init()
	return ske
}
