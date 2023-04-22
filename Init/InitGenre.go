package Init

import "github.com/pritamdas99/Movie_Server/Data"

type genreDB map[int]Data.Genre

var GenreList genreDB

func init() {
	GenreList = make(genreDB)
	GenreList[1] = Data.Genre{
		Name: "Romantic",
		Id:   1,
	}
	GenreList[2] = Data.Genre{
		Name: "Action",
		Id:   2,
	}
	GenreList[3] = Data.Genre{
		Name: "Comedy",
		Id:   3,
	}
}
