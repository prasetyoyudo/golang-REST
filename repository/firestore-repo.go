package repository

import (
	"context"
	"log"

	"../entity"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type repo struct {}

//NewFirestore Repository
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectId 		string = "blog-api-c8dda"
	collectionName 	string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId )
	if err != nil {
		log.Fatalf("failed to create a firestore client: %v", err)
		return nil, err
	}

	defer client.Close()
	client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":	post.ID,
		"Title":	post.Title,
		"Description":	post.Description,
	})

	if err != nil {
		log.Fatalln("failed adding a new post %v", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId )
	if err != nil {
		log.Fatalf("failed to create a firestore client: %v", err)
		return nil, err
	}	
	defer client.Close()
	var posts []entity.Post
	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
    
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed to iterate through collection: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Description:  doc.Data()["Description"].(string),
		}
		posts = append(posts,post)
	}
	return posts , nil 
}
