package main

import (
	"fmt"
	"github.com/pritamdas99/Movie_Server/cmd"
)

//var router = APIs.Router
//
//func callFunc() {
//	router.HandleFunc("/", APIs.Homepage)
//	router.HandleFunc("/adddirector", APIs.AddDirector).Methods("POST")
//	router.HandleFunc("/deletedirector", APIs.DeleteDirector).Methods("DELETE")
//	router.HandleFunc("/getdirector", APIs.GetDirector).Methods("GET")
//	router.HandleFunc("/updatedirector", APIs.UpdateDirector).Methods("PUT")
//
//	router.HandleFunc("/addgenre", APIs.AddGenre).Methods("POST")
//	router.HandleFunc("/deletegenre", APIs.DeleteGenre).Methods("DELETE")
//	router.HandleFunc("/getgenre", APIs.GetGenre).Methods("GET")
//	router.HandleFunc("/updategenre", APIs.UpdateGenre).Methods("PUT")
//
//	router.HandleFunc("/adduser", APIs.AddUser).Methods("POST")
//	router.HandleFunc("/deleteuser", APIs.DeleteUser).Methods("DELETE")
//	router.HandleFunc("/getuser", APIs.GetUser).Methods("GET")
//	router.HandleFunc("/updateuser", APIs.UpdateUser).Methods("PUT")
//
//	router.HandleFunc("/addmovie", APIs.AddMovie).Methods("POST")
//	router.HandleFunc("/deletemovie", APIs.DeleteMovie).Methods("DELETE")
//	router.HandleFunc("/deletemoviegenre/{id}", APIs.DeleteMovieGenre).Methods("DELETE")
//	router.HandleFunc("/addmoviegenre/{id}", APIs.AddMovieGenre).Methods("GET")
//	router.HandleFunc("/deletemoviedirector/{id}", APIs.DeleteMovieDirector).Methods("DELETE")
//	router.HandleFunc("/addmoviedirector/{id}", APIs.AddMovieDirector).Methods("GET")
//	router.HandleFunc("/getmovie", APIs.GetMovie).Methods("GET")
//	router.HandleFunc("/getmoviebygenre", APIs.GetMovieByGenre).Methods("GET")
//	router.HandleFunc("/getmoviebydirector", APIs.GetMovieByDirector).Methods("GET")
//	router.HandleFunc("/updatemovie", APIs.UpdateMovie).Methods("PUT")
//}
//
//func StartServer(port int) {
//	log.Printf("-------------starting sever at %d -------\n", port)
//	callFunc()
//	Server := &http.Server{
//		Addr:    ":" + strconv.Itoa(port),
//		Handler: router,
//	}
//
//	fmt.Println(Server.ListenAndServe())
//
//}

func main() {
	fmt.Println("search")
	cmd.Execute()
}
