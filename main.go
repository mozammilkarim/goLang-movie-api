package main
import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}
type Director struct{
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}
var movies []Movie

func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}
func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item:=range movies{
		if item.ID==params["id"]{
			movies=append(movies[:index],movies[index+1:]...)
			break
		}
	}
}
func getMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for _,item:=range movies{
		if item.ID==params["id"]{
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}
func createMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie;
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID=strconv.Itoa(rand.Intn(100000000))
	movies=append(movies,movie)
	json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item:=range movies{
		if item.ID==params["id"]{
			movies=append(movies[:index],movies[index+1:]...)
			var movie Movie;
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID=params["id"]
			movies=append(movies,movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
func main(){
	r:=mux.NewRouter()
	movies =append(movies,Movie{ID:"1",Isbn:"129090",Title:"Bloody Moonday",Director:&Director{FirstName:"Michael",LastName:"jackson"}})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	fmt.Printf("The Server is starting at port 8000")
	log.Fatal(http.ListenAndServe(":8000",r))
}