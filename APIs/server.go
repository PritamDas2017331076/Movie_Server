package APIs

import (
	"fmt"
	"github.com/pritamdas99/Movie_Server/Auth"
	"log"
	"net/http"
	"strconv"
)

func callFunc() {
	router.HandleFunc("/", Auth.IsAuthenticated(Homepage))
	router.HandleFunc("/adddirector", Auth.IsAuthenticated(AddDirector)).Methods("POST")
	router.HandleFunc("/deletedirector", Auth.IsAuthenticated(DeleteDirector)).Methods("DELETE")
	router.HandleFunc("/getdirector", Auth.IsAuthenticated(GetDirector)).Methods("GET")
	router.HandleFunc("/updatedirector/{id}", Auth.IsAuthenticated(UpdateDirector)).Methods("PUT")

	router.HandleFunc("/addgenre", Auth.IsAuthenticated(AddGenre)).Methods("POST")
	router.HandleFunc("/deletegenre", Auth.IsAuthenticated(DeleteGenre)).Methods("DELETE")
	router.HandleFunc("/getgenre", Auth.IsAuthenticated(GetGenre)).Methods("GET")
	router.HandleFunc("/updategenre/{id}", Auth.IsAuthenticated(UpdateGenre)).Methods("PUT")

	router.HandleFunc("/adduser", AddUser).Methods("POST")
	router.HandleFunc("/deleteuser", Auth.IsAuthenticated(DeleteUser)).Methods("DELETE")
	router.HandleFunc("/getuser", GetUser).Methods("GET")
	router.HandleFunc("/updateuser/{id}", Auth.IsAuthenticated(UpdateUser)).Methods("PUT")
	router.HandleFunc("/loginuser", LoginUser).Methods("POST")
	router.HandleFunc("/logout", Logout).Methods("GET")

	router.HandleFunc("/addmovie", Auth.IsAuthenticated(AddMovie)).Methods("POST")
	router.HandleFunc("/deletemovie", Auth.IsAuthenticated(DeleteMovie)).Methods("DELETE")
	router.HandleFunc("/deletemoviegenre/{id}", Auth.IsAuthenticated(DeleteMovieGenre)).Methods("DELETE")
	router.HandleFunc("/addmoviegenre/{id}", Auth.IsAuthenticated(AddMovieGenre)).Methods("GET")
	router.HandleFunc("/deletemoviedirector/{id}", Auth.IsAuthenticated(DeleteMovieDirector)).Methods("DELETE")
	router.HandleFunc("/addmoviedirector/{id}", Auth.IsAuthenticated(AddMovieDirector)).Methods("GET")
	router.HandleFunc("/getmovie", Auth.IsAuthenticated(GetMovie)).Methods("GET")
	router.HandleFunc("/getmoviebygenre", Auth.IsAuthenticated(GetMovieByGenre)).Methods("GET")
	router.HandleFunc("/getmoviebydirector", Auth.IsAuthenticated(GetMovieByDirector)).Methods("GET")
	router.HandleFunc("/updatemovie/{id}", Auth.IsAuthenticated(UpdateMovie)).Methods("PUT")
}

func StartServer(port int) {
	log.Printf("-------------starting sever at %d -------\n", port)
	callFunc()
	Server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: router,
	}

	fmt.Println(Server.ListenAndServe())

}
