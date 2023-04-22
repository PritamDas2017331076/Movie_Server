package APIs

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pritamdas99/Movie_Server/Data"
	"github.com/pritamdas99/Movie_Server/Init"
	"net/http"
	"strconv"
)

func AddGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	avail := 1
	for idx, _ := range Init.GenreList {
		if idx > avail {
			avail = idx
		}
	}
	var NewGenre Data.Genre
	err := json.NewDecoder(r.Body).Decode(&NewGenre)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	Init.GenreList[avail+1] = NewGenre
	err = json.NewEncoder(w).Encode(Init.GenreList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	ids := params.Get("id")
	idi, _ := strconv.Atoi(ids)
	delete(Init.GenreList, idi)
	err := json.NewEncoder(w).Encode(Init.GenreList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mk []Data.Genre
	for _, genre := range Init.GenreList {
		mk = append(mk, genre)
	}
	err := json.NewEncoder(w).Encode(mk)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func UpdateGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ids := params["id"]
	idi, _ := strconv.Atoi(ids)
	Genre := Init.GenreList[idi]
	err := json.NewDecoder(r.Body).Decode(&Genre)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	Init.GenreList[idi] = Genre
	err = json.NewEncoder(w).Encode(Genre)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}
