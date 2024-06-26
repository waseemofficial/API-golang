package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/waseemofficial/API_golang/entity"
	"google.golang.org/api/iterator"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
type repo struct{}

const (
	projectId      string = "api-golang-22fab"
	collectionName string = "posts"
)

// NeewPostRepositroy
func NewPostRepository() PostRepository {
	return &repo{}
}

// Save implements PostRepository.
func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatalf("Failed adding a new post:%v", err)
		return nil, err
	}
	return post, nil
}

// FindAll implements PostRepository.
func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("failed to create firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of Posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Text:  doc.Data()["Text"].(string),
			Title: doc.Data()["Title"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
