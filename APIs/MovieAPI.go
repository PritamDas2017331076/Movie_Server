package APIs

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pritamdas99/Movie_Server/Data"
	"github.com/pritamdas99/Movie_Server/Init"
	"net/http"
	"strconv"
)

type ddd struct {
	Did int `json:"did"`
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	avail := 1
	for idx, _ := range Init.MovieList {
		if idx > avail {
			avail = idx
		}
	}
	var NewMovie Data.Movie
	err := json.NewDecoder(r.Body).Decode(&NewMovie)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	Init.MovieList[avail+1] = NewMovie
	err = json.NewEncoder(w).Encode(Init.MovieList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	ids := params.Get("id")
	idi, _ := strconv.Atoi(ids)
	delete(Init.MovieList, idi)
	err := json.NewEncoder(w).Encode(Init.MovieList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func DeleteMovieGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ids := params["id"]
	idi, _ := strconv.Atoi(ids)
	var data ddd
	Movie := Init.MovieList[idi]
	fmt.Println("init", data)
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("got an error")
		fmt.Println(err.Error(), err)
	}
	fmt.Println(data, data.Did, r.Body, data, data.Did)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	ind := -1
	fmt.Println("ind", ind, "idi", idi)
	for idi, val := range Movie.Genres {
		if val == data.Did {
			ind = idi
			break
		}
	}
	fmt.Println(ind)
	if ind != -1 {
		lst := Movie.Genres
		lst = append(lst[:ind], lst[ind+1:]...)
		Movie.Genres = lst
		Init.MovieList[idi] = Movie
	}
	err = json.NewEncoder(w).Encode(Init.MovieList[idi])
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}

func AddMovieGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ids := params["id"]
	idi, _ := strconv.Atoi(ids)

	var data ddd
	Movie := Init.MovieList[idi]
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	val := Movie.Genres
	fl := 0
	for _, dir := range val {
		if dir == data.Did {
			fl = 1
			break
		}
	}
	if fl == 0 {
		lst := append(val, data.Did)
		Movie.Genres = lst
		Init.MovieList[idi] = Movie
	}
	err = json.NewEncoder(w).Encode(Init.MovieList[idi])
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func DeleteMovieDirector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ids := params["id"]
	idi, _ := strconv.Atoi(ids)
	var data ddd
	Movie := Init.MovieList[idi]
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	ind := -1
	for idi, val := range Movie.Directors {
		if val == data.Did {
			ind = idi
			break
		}
	}
	if ind != -1 {
		lst := Movie.Directors
		lst = append(lst[:ind], lst[ind+1:]...)
		Movie.Directors = lst
		Init.MovieList[idi] = Movie
	}
	err = json.NewEncoder(w).Encode(Init.MovieList[idi])
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}

func AddMovieDirector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ids := params["id"]
	idi, _ := strconv.Atoi(ids)
	var data ddd
	Movie := Init.MovieList[idi]
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	val := Movie.Directors
	fl := 0
	for _, dir := range val {
		if dir == data.Did {
			fl = 1
			break
		}
	}
	if fl == 0 {
		lst := append(val, data.Did)
		Movie.Directors = lst
		Init.MovieList[idi] = Movie
	}
	err = json.NewEncoder(w).Encode(Init.MovieList[idi])
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mk []Data.Movie
	for _, movie := range Init.MovieList {
		mk = append(mk, movie)
	}
	err := json.NewEncoder(w).Encode(mk)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func GetMovieByGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	ids := params.Get("id")
	idi, _ := strconv.Atoi(ids)
	fmt.Println("idi", idi)
	var mk []Data.Movie
	for _, movie := range Init.MovieList {
		fl := 0
		for _, ip := range movie.Genres {
			if ip == idi {
				fl = 1
				break
			}
		}
		if fl == 1 {
			mk = append(mk, movie)
		}
	}
	err := json.NewEncoder(w).Encode(mk)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func GetMovieByDirector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	ids := params.Get("id")
	idi, _ := strconv.Atoi(ids)
	var mk []Data.Movie
	for _, movie := range Init.MovieList {
		fl := 0
		for _, ip := range movie.Directors {
			if ip == idi {
				fl = 1
				break
			}
		}
		if fl == 1 {
			mk = append(mk, movie)
		}
	}
	err := json.NewEncoder(w).Encode(mk)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ids := params["id"]
	idi, _ := strconv.Atoi(ids)
	Movie := Init.MovieList[idi]
	err := json.NewDecoder(r.Body).Decode(&Movie)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = json.NewEncoder(w).Encode(Movie)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}
