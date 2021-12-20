package mongodb

type Token struct {
	Id    string  `bson:"_id,omitempty"`
	Name  string  `bson:"name,omitempty"`
	Price float64 `bson:"price,omitempty"`
}

type Votes struct {
	Id string `bson:"_id,omitempty"`
	To string `bson:"name,omitempty"`
}
