package main

import (
	"github.com/wuyutaott/leaf"
	leafconf "github.com/wuyutaott/leaf/conf"
	"github.com/wuyutaott/leaf/test/xxx_server/xxx_conf"
	"github.com/wuyutaott/leaf/test/xxx_server/xxx_game"
	"github.com/wuyutaott/leaf/test/xxx_server/xxx_gate"
)

func main() {
	leafconf.LogLevel = xxx_conf.Server.LogLevel
	leafconf.LogPath = xxx_conf.Server.LogPath
	leafconf.LogFlag = xxx_conf.LogFlag
	leafconf.ConsolePort = xxx_conf.Server.ConsolePort
	leafconf.ProfilePath = xxx_conf.Server.ProfilePath

	leaf.Run(
		xxx_game.Module,
		xxx_gate.Module,
	)
}
