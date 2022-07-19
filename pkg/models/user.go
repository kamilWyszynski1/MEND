package models

// User represents user's data.
// NOTE: we could use some validation here but it is skipped for simplicity
// we just assume that user will put proper data as input.
// For validation purposes we could use this lib: https://github.com/go-playground/validator.
type User struct {
	// NOTE: ID field could be separated from User in order no to
	// get this during PUT but it skipped for simplicity.
	ID      int    `json:"id" bson:"_id"`
	Name    string `json:"name" bson:"name"`
	Surname string `json:"surname" bson:"surname"`
}
