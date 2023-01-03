package xxx_conf

import (
	"encoding/json"
	"github.com/wuyutaott/leaf/log"
	"os"
)

var Server struct {
	LogLevel        string
	LogPath         string
	WSAddr          string
	CertFile        string
	KeyFile         string
	TCPAddr         string
	MaxConnNum      int
	ConsolePort     int
	ProfilePath     string
	MgodbAddr       string
	LoginMgoConnNum int
	GameMgoConnNum  int
}

func init() {
	data, err := os.ReadFile("xxx_conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
