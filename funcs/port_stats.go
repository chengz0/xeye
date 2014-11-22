package funcs

import (
	"github.com/chengz0/xstat/collector"
	"github.com/ulricqin/goutils/slicetool"
)

func PortIsListen(port int64) bool {
	return slicetool.SliceContainsInt64(collector.ListenPorts(), port)
}
