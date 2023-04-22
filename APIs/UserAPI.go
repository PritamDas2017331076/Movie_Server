package APIs

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pritamdas99/Movie_Server/Auth"
	"github.com/pritamdas99/Movie_Server/Data"
	"github.com/pritamdas99/Movie_Server/Init"
	"log"
	"net/http"
	"os"
	"time"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var NewUser Data.User
	err := json.NewDecoder(r.Body).Decode(&NewUser)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	Init.UserList[NewUser.Id] = NewUser
	err = json.NewEncoder(w).Encode(Init.UserList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	ids := params.Get("id")
	delete(Init.UserList, ids)
	err := json.NewEncoder(w).Encode(Init.UserList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mk []Data.User
	for _, user := range Init.UserList {
		mk = append(mk, user)
	}
	err := json.NewEncoder(w).Encode(mk)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ids := params["id"]
	User := Init.UserList[ids]
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	Init.UserList[ids] = User
	err = json.NewEncoder(w).Encode(User)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliation/json")
	var NewUser Data.User
	err := json.NewDecoder(r.Body).Decode(&NewUser)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	val, ok := Init.UserList[NewUser.Id]
	if !ok {
		http.Error(w, err.Error(), 400)
		return
	}
	if val.Password != NewUser.Password {
		http.Error(w, err.Error(), 400)
		return
	}
	token, err := Auth.GenerateJWT()
	if err != nil {
		log.Fatal(err)
	}
	cookie := http.Cookie{
		Name:    "Token",
		Value:   token,
		Path:    "/",
		Expires: time.Now().Add(15 * time.Minute),
	}
	os.Setenv("Token", token)
	fmt.Println("gotenv token", os.Getenv("Token"))
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(NewUser)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Expires: time.Now(),
	})
	w.WriteHeader(http.StatusOK)
}

// curl -H "Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2ODIyMDAyNjcsInVzZXIiOiJwcml0YW0ifQ.l3uYLQkQJD6Qk_TRcdJZ_FXbBUOabzdDNBfQNBr1Tds" GET  http://localhost:8080/getmovie
