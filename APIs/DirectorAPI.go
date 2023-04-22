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

var router = mux.NewRouter()

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func AddDirector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	avail := 1
	for idx, _ := range Init.DirectorList {
		//fmt.Println(idx, val, avail)
		if idx > avail {
			avail = idx
		}
	}
	var NewDirector Data.Director
	err := json.NewDecoder(r.Body).Decode(&NewDirector)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	Init.DirectorList[avail+1] = NewDirector
	err = json.NewEncoder(w).Encode(Init.DirectorList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteDirector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	ids := params.Get("id")
	idi, _ := strconv.Atoi(ids)
	delete(Init.DirectorList, idi)
	err := json.NewEncoder(w).Encode(Init.DirectorList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetDirector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mk []Data.Director
	for _, director := range Init.DirectorList {
		mk = append(mk, director)
	}
	err := json.NewEncoder(w).Encode(mk)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func UpdateDirector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ids := params["id"]
	idi, _ := strconv.Atoi(ids)
	Director := Init.DirectorList[idi]
	fmt.Println(Director)
	err := json.NewDecoder(r.Body).Decode(&Director)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	Init.DirectorList[idi] = Director
	err = json.NewEncoder(w).Encode(Director)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}
