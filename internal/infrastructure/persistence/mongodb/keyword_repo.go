package mongodb

import (
	"context"
	"time"

	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type KeywordRepo struct {
	DB *mongo.Database
}

func (repo *KeywordRepo) SaveNewKeyword(keyword string) error {
	coll := repo.DB.Collection("keyword_histories")
	now := time.Now()
	doc := &entity.KeywordHistory{
		Keyword:   keyword,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}

	return nil
}

func (repo *KeywordRepo) GetPopularKeyword() ([]entity.Keyword, error) {
	return nil, nil
}
