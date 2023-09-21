package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	Id    int    `json:"id"`
	Text  string `json:"text"`
	Title string `json:"title"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{{Id: 1, Title: "title 1", Text: "text 1"}}
}

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "hi")
	})
	router.HandleFunc("/posts", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-type", "application/json")
		fmt.Println("posts")
		result, err := json.Marshal(posts)
		if err != nil {
			fmt.Println(err)
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(`{"error":"Error marshalling the posts array"}`))
			return
		}
		resp.WriteHeader(http.StatusOK)
		resp.Write(result)
	}).Methods("GET")
	router.HandleFunc("/post", func(resp http.ResponseWriter, req *http.Request) {
		var post Post
		fmt.Println("post")
		err := json.NewDecoder(req.Body).Decode(&post)
		if err != nil {
			fmt.Println(err)
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(`{"error":"Error unmarshalling the request"}`))
			return
		}
		post.Id = len(posts) + 1
		posts = append(posts, post)
		resp.WriteHeader(http.StatusOK)
		result, err := json.Marshal(post)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(`{"error":"Error unmarshalling the request"}`))
			return
		}
		resp.Write(result)

	}).Methods("POST")
	log.Println("hi1")
	http.ListenAndServe(port, router)
}
