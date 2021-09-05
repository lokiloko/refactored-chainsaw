package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/dto"
	"net/http"
	"os"
	"runtime"
	"time"
)

func (controller *Controller) Health(c echo.Context) error {
	var (
		response dto.HealthResponse
		m        runtime.MemStats
	)

	response.Pid = os.Getpid()
	response.Uptime = time.Since(time.Now())
	response.Status = "ok"
	runtime.ReadMemStats(&m)
	response.Memory.Alloc = m.Alloc
	response.Memory.TotalAlloc = m.TotalAlloc
	response.Memory.Sys = m.Sys
	response.Memory.NumGC = m.NumGC
	response.Memory.HeapAlloc = m.HeapAlloc
	response.Memory.HeapSys = m.HeapSys
	response.NumCPU = runtime.NumCPU()

	return c.JSON(http.StatusOK, response)
}
