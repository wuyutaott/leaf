package internal

import (
	"github.com/wuyutaott/leaf/gate"
	"github.com/wuyutaott/leaf/test/xxx_server/xxx_conf"
	"github.com/wuyutaott/leaf/test/xxx_server/xxx_game"
	"github.com/wuyutaott/leaf/test/xxx_server/xxx_msg"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      xxx_conf.Server.MaxConnNum,
		PendingWriteNum: xxx_conf.PendingWriteNum,
		MaxMsgLen:       xxx_conf.MaxMsgLen,
		WSAddr:          xxx_conf.Server.WSAddr,
		HTTPTimeout:     xxx_conf.HTTPTimeout,
		CertFile:        xxx_conf.Server.CertFile,
		KeyFile:         xxx_conf.Server.KeyFile,
		TCPAddr:         xxx_conf.Server.TCPAddr,
		LenMsgLen:       xxx_conf.LenMsgLen,
		LittleEndian:    xxx_conf.LittleEndian,
		Processor:       xxx_msg.Processor,
		AgentChanRPC:    xxx_game.ChanRPC,
	}
}
