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
	GetByIMDBIDResponse struct {
		Data Movie `json:"data"`
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
	Movie struct {
		Title      string `json:"title"`
		Year       string `json:"year"`
		Rated      string `json:"rated"`
		Released   string `json:"released"`
		Runtime    string `json:"runtime"`
		Genre      string `json:"genre"`
		Director   string `json:"director"`
		Writer     string `json:"writer"`
		Actors     string `json:"actors"`
		Plot       string `json:"plot"`
		Language   string `json:"language"`
		Country    string `json:"country"`
		Awards     string `json:"awards"`
		Poster     string `json:"poster"`
		Metascore  string `json:"metascore"`
		ImdbID     string `json:"imdbID"`
		ImdbVotes  string `json:"imdbVotes"`
		ImdbRating string `json:"imdbRating"`
		Type       string `json:"type"`
		DVD        string `json:"dvd"`
		Production string `json:"production"`
		Website    string `json:"website"`
		BoxOffice  string `json:"boxOffice"`
	}
)
