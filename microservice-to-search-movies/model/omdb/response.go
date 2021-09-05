package omdb

type (
	GetPaginatedResponse struct {
		Search      []Movie `json:"Search"`
		TotalResult uint64  `json:"totalResults"`
		Response    string  `json:"Response"`
	}
	GetByIDResponse struct {
		Movie
		Response string `json:"Response"`
	}
	Movie struct {
		Title      string `json:"Title"`
		Year       string `json:"Year"`
		Rated      string `json:"Rated"`
		Released   string `json:"Released"`
		Runtime    string `json:"Runtime"`
		Genre      string `json:"Genre"`
		Director   string `json:"Director"`
		Writer     string `json:"Writer"`
		Actors     string `json:"Actors"`
		Plot       string `json:"Plot"`
		Language   string `json:"Language"`
		Country    string `json:"Country"`
		Awards     string `json:"Awards"`
		Poster     string `json:"Poster"`
		Metascore  string `json:"Metascore"`
		ImdbID     string `json:"imdbID"`
		ImdbVotes  string `json:"imdbVotes"`
		ImdbRating string `json:"imdbRating"`
		Type       string `json:"Type"`
		DVD        string `json:"DVD"`
		Production string `json:"Production"`
		Website    string `json:"Website"`
		BoxOffice  string `json:"BoxOffice"`
	}
)
