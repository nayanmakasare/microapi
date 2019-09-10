package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path"
	"strings"
)


func main(){
	fmt.Println("Starting server program")
	fmt.Println("http://localhost:12345/cloudwalkerPrimePages")
	router := mux.NewRouter()
	router.HandleFunc("/cloudwalkerPrimePages", GetPrimeUrl).Methods("GET")
	router.HandleFunc("/{key}", Pages).Methods("GET")
	router.Path("/{key}/{subkey}").HandlerFunc(GetPageData).Methods("GET")
	err := http.ListenAndServe(":12345", router)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("End and exit")
}

func Pages(response http.ResponseWriter, request *http.Request){
	response.Header().Set("content-type", "application/json")
	vars := mux.Vars(request)
	fmt.Println("********************************")
	fmt.Printf("http://localhost:12345/%s/%s\n",vars["key"],"carousel")
	fmt.Printf("http://localhost:12345/%s/%s\n",vars["key"],"rowContent1")
	fmt.Printf("http://localhost:12345/%s/%s\n",vars["key"],"rowContent2")
	fmt.Printf("http://localhost:12345/%s/%s\n",vars["key"],"rowContent3")
	fmt.Printf("http://localhost:12345/%s/%s\n",vars["key"],"rowContent4")
	fmt.Printf("http://localhost:12345/%s/%s\n",vars["key"],"rowContent5")
	response.WriteHeader(http.StatusOK)
	filePath := path.Join(".", vars["key"]+".json")
	http.ServeFile(response, request, filePath)
}

func GetPageData(response http.ResponseWriter, request *http.Request)  {
	vars := mux.Vars(request)
	if(!strings.Contains(request.RequestURI, "carousel")){
		fmt.Println("************************************************")
		fmt.Println("To Fetch Full MovieTile Content")
		fmt.Println("http://localhost:12345/5d314b25e2f82960caff600c")
	}
	response.WriteHeader(http.StatusOK)
	filePath := path.Join(vars["key"], vars["subkey"]+".json")
	http.ServeFile(response, request, filePath)
}

func GetPrimeUrl(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	response.WriteHeader(http.StatusOK)
	filePath := path.Join(".", "cloudwalkerPrimePages.json")
	if(!strings.Contains(request.RequestURI, "5d314b25e2f82960caff600c")){
		fmt.Println("************************************************")
		fmt.Println("http://localhost:12345/home")
		fmt.Println("http://localhost:12345/movie")
		fmt.Println("http://localhost:12345/tvshows")
		fmt.Println("http://localhost:12345/more")
	}
	http.ServeFile(response, request, filePath)
}






















