package product

import (
	"context"
	"fmt"
	"github.com/djairdj/golang-desafio-tecnico1/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) Repository {
	collectionProduct := db.Collection("product")
	return &mongoRepository{
		collection: collectionProduct,
	}
}
func (r mongoRepository) Create(ctx context.Context, name string) (*entity.Product, error) {
	data := bson.D{{"name", name}, {"votes", 0}}
	res, err := r.collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	id, ok := res.InsertedID.(primitive.ObjectID) //feitura de cast
	if !ok {
		return nil, fmt.Errorf("nao foi possivel converter o id")
	}

	p := entity.Product{
		ID:    id.Hex(),
		Name:  name,
		Votes: 0,
	}

	return &p, nil
}

func (r mongoRepository) List(ctx context.Context) ([]entity.Product, error) {
	findOptions := options.Find()
	findOptions.SetLimit(8)
	results := []entity.Product{}

	cur, err := r.collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("could not find the products: %v", err)
	}

	for cur.Next(ctx) {
		//Create a value into which the single document can be decoded
		elem := entity.Product{}
		err = cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	err = cur.Close(ctx)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r mongoRepository) GetOne(ctx context.Context, id string) (*entity.Product, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := r.collection.FindOne(ctx, bson.M{"_id": objectId})
	if result.Err() != nil {
		return nil, result.Err()
	}

	prod := entity.Product{}
	err = result.Decode(&prod)
	if err != nil {
		return nil, err
	}

	return &prod, nil
}

func (r mongoRepository) Update(ctx context.Context, product *entity.Product) error {
	objectId, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}

	update := bson.M{
		"$set": bson.M{
			"name":  product.Name,
			"votes": product.Votes,
		},
	}

	upsert := options.FindOneAndUpdate().SetUpsert(true)
	single := r.collection.FindOneAndUpdate(ctx, filter, update, upsert)
	if single.Err() != nil {
		return single.Err()
	}

	return nil
}
