package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/waseemofficial/API_golang/entity"
	"github.com/waseemofficial/API_golang/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "hi")
	})
	router.HandleFunc("/posts", func(response http.ResponseWriter, req *http.Request) {
		response.Header().Set("Content-type", "application/json")
		fmt.Println("posts")
		posts, err := repo.FindAll()
		if err != nil {
			fmt.Println(err)
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"error":"Error getting the posts []"}`))
			return
		}

		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(posts)
	}).Methods("GET")

	router.HandleFunc("/post", func(response http.ResponseWriter, req *http.Request) {
		var post entity.Post
		fmt.Println("post")
		err := json.NewDecoder(req.Body).Decode(&post)
		if err != nil {
			fmt.Println(err)
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"error":"Error unmarshalling the request"}`))
			return
		}
		post.ID = rand.Int63()
		repo.Save(&post)
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(post)

	}).Methods("POST")
	log.Println(port, router)
	http.ListenAndServe(port, router)
}
