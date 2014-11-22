package http

import (
	"github.com/chengz0/xstat/collector"
	// "io"
	// "log"
	"net/http"
)

func LoginRouter() {
	m.Get("/proc/lastlogin", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		outs, err := collector.LastLogin()
		if err != nil {
			return RenderErrDto(err.Error())
		}

		var ret [][]interface{} = make([][]interface{}, 0)
		for _, out := range outs {
			ret = append(ret, []interface{}{
				out.Username,
				out.From,
				out.Latest,
			})
		}
		return RenderDataDto(ret)
	})
}
