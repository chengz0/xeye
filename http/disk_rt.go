package http

import (
	"fmt"
	"github.com/chengz0/xstat/collector"
	"github.com/ulricqin/goutils/formatter"
	// "log"
	"net/http"
)

func DiskRouter() {
	m.Get("/proc/df/bytes", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		mountPoints, err := collector.ListMountPoint()
		if err != nil {
			return RenderErrDto(err.Error())
		}

		var ret [][]interface{} = make([][]interface{}, 0)
		for idx := range mountPoints {
			var du *collector.DeviceUsageStruct
			du, err = collector.BuildDeviceUsage(mountPoints[idx])
			if err == nil {
				ret = append(ret,
					[]interface{}{
						du.FsSpec,
						formatter.DisplaySize(float64(du.BlocksAll)),
						formatter.DisplaySize(float64(du.BlocksUsed)),
						formatter.DisplaySize(float64(du.BlocksFree)),
						fmt.Sprintf("%.1f%%", du.BlocksUsedPercent),
						du.FsFile,
						formatter.DisplaySize(float64(du.InodesAll)),
						formatter.DisplaySize(float64(du.InodesUsed)),
						formatter.DisplaySize(float64(du.InodesFree)),
						fmt.Sprintf("%.1f%%", du.BlocksUsedPercent),
						du.FsVfstype,
					})
			}
		}
		return RenderDataDto(ret)
	})
}

func SdaRouter() {
	m.Get("/proc/sda", func(w http.ResponseWriter) string {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		sda, err := collector.GetCurrentDisk()
		if err != nil {
			return RenderErrDto(err.Error())
		}

		return RenderDataDto([]interface{}{sda.Total / 1024 / 1024, sda.Used / 1024 / 1024, sda.Free / 1024 / 1024})
	})
}
