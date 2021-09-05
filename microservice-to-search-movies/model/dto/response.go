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
	GetMoviesPaginatedResponse struct {
		Data Movies `json:"data"`
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
	Movies []Movie
	Movie struct {
		Title      string  `json:"title,omitempty"`
		Year       string  `json:"year,omitempty"`
		Rated      *string `json:"rated,omitempty"`
		Released   *string `json:"released,omitempty"`
		Runtime    *string `json:"runtime,omitempty"`
		Genre      *string `json:"genre,omitempty"`
		Director   *string `json:"director,omitempty"`
		Writer     *string `json:"writer,omitempty"`
		Actors     *string `json:"actors,omitempty"`
		Plot       *string `json:"plot,omitempty"`
		Language   *string `json:"language,omitempty"`
		Country    *string `json:"country,omitempty"`
		Awards     *string `json:"awards,omitempty"`
		Poster     string  `json:"poster,omitempty"`
		Metascore  *string `json:"metascore,omitempty"`
		ImdbID     string  `json:"imdbID,omitempty"`
		ImdbVotes  *string `json:"imdbVotes,omitempty"`
		ImdbRating *string `json:"imdbRating,omitempty"`
		Type       string  `json:"type,omitempty"`
		DVD        *string `json:"dvd,omitempty"`
		Production *string `json:"production,omitempty"`
		Website    *string `json:"website,omitempty"`
		BoxOffice  *string `json:"boxOffice,omitempty"`
	}
)
