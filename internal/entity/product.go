package entity

type Product struct {
	ID    string `bson:"_id"`
	Name  string `bson:"name"`
	Votes int32  `bson:"votes"`
}
