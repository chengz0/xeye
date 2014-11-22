// render.go
package main

import (
	"github.com/chengz0/xeye/funcs"
	"github.com/chengz0/xeye/http"
	"log"
	"time"
)

func main() {

	go InitData()

	log.Println("servering ...")
	http.StartHttpServer()
}

func InitData() {
	for {
		funcs.UpdateCpuStat()
		funcs.UpdateDiskStats()
		funcs.UpdateIfStat()

		time.Sleep(time.Duration(1) * time.Second)
	}
}
