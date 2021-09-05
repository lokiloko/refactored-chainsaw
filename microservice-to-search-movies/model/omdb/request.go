package omdb

type (
	GetPaginatedRequest struct {
		Keyword string `json:"keyword"`
		Page    uint64 `json:"page"`
	}
	GetByIDRequest struct {
		ID string `json:"imdbID"`
	}
)
