package dto

import "time"

type (
	HealthResponse struct {
		Pid    int           `json:"pid"`
		Uptime time.Duration `json:"uptime"`
		Memory MemoryUsage   `json:"memory"`
		NumCPU int           `json:"numCPU"`
		Status string        `json:"status"`
	}
)

type (
	MemoryUsage struct {
		Alloc      uint64 `json:"alloc"`
		TotalAlloc uint64 `json:"totalAlloc"`
		Sys        uint64 `json:"sys"`
		HeapAlloc  uint64 `json:"heapAlloc"`
		HeapSys    uint64 `json:"heapSys"`
		NumGC      uint32 `json:"numGC"`
	}
)
