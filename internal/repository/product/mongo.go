package product

import "go.mongodb.org/mongo-driver/mongo"

type mongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) Repository {
	collectionProduct := db.Collection("product")
	return &mongoRepository{
		collection: collectionProduct,
	}
}
