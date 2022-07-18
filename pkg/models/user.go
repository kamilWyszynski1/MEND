package models

// User represents user's data.
type User struct {
	ID      int    `json:"id" bson:"_id"`
	Name    string `json:"name" bson:"name"`
	Surname string `json:"surname" bson:"surname"`
}
