package Init

import "github.com/pritamdas99/Movie_Server/Data"

type directorDB map[int]Data.Director

var DirectorList directorDB

func init() {
	DirectorList = make(directorDB)
	DirectorList[1] = Data.Director{
		Name: "Nolan",
		Id:   1,
	}
	DirectorList[2] = Data.Director{
		Name: "Rowling",
		Id:   2,
	}
	DirectorList[3] = Data.Director{
		Name: "Dobby",
		Id:   3,
	}
}
