package repository

import (
	"../entity"
	"cloud.google.com/go/firestore"
)

type PostRepository interface {
	Save(post *entity.Post) {*entity.Post, error}
	FindAll() ([]entity.Post, error) 
}

type repo struct {}

//NewPoostRepo
func NewPoostRepo() PostRepository {
	return &repo{}
}

func (*repo) Save(post *entity.Post) {*entity.Post, error} {
	
}
