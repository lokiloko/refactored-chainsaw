package mocks

import (
	_ "embed"
	"encoding/json"
	"github.com/lokiloko/refactored-chainsaw/microservice-to-search-movies/model/omdb"
)

//go:embed service/omdb/data/getById.json
var getByIdString1 string

//go:embed service/omdb/data/getPaginated.json
var getPaginatedString1 string


func GetOMDBByID(id uint64) *omdb.GetByIDResponse {
	var (
		data       omdb.GetByIDResponse
		dataString string
	)

	switch id {
	case 1:
		dataString = getByIdString1
	}

	err := json.Unmarshal([]byte(dataString), &data)

	if err != nil {
		panic("JSON Unmarshal failed, " + err.Error())
	}

	return &data
}

func GetOmdbPaginated(id uint64) *omdb.GetPaginatedResponse {
	var (
		data       omdb.GetPaginatedResponse
		dataString string
	)

	switch id {
	case 1:
		dataString = getPaginatedString1
	}

	err := json.Unmarshal([]byte(dataString), &data)

	if err != nil {
		panic("JSON Unmarshal failed, " + err.Error())
	}

	return &data
}
