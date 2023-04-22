package Init

import "github.com/pritamdas99/Movie_Server/Data"

type movieDB map[int]Data.Movie

var MovieList movieDB

func init() {
	MovieList = make(movieDB)
	MovieList[1] = Data.Movie{
		Name: "Twilight",
		Id:   1,
		Directors: []int{
			1, 2,
		},
		Genres: []int{
			2, 3,
		},
	}
	MovieList[2] = Data.Movie{
		Name: "With Love From Paris",
		Id:   2,
		Directors: []int{
			2, 3,
		},
		Genres: []int{
			1, 2,
		},
	}
	MovieList[3] = Data.Movie{
		Name: "With Love From UK",
		Id:   3,
		Directors: []int{
			2, 3,
		},
		Genres: []int{
			1, 2, 3,
		},
	}
}
