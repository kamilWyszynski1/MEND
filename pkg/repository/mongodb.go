package repository

import (
	"MEND/pkg/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Mongo is a handler for MongoDB operations for user entity.
type Mongo struct {
	client *mongo.Collection // 'user' collection.
}

// Create inserts new user.
func (m *Mongo) Create(ctx context.Context, user models.User) error {
	_, err := m.client.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to insert user, %w", err)
	}
	return nil
}

// Get finds user with given id.
func (m *Mongo) Get(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	if err := m.client.FindOne(ctx, bson.D{{Key: "id", Value: id}}).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to find user, %w", err)
	}
	return &user, nil
}

// Update updates user entry.
func (m *Mongo) Update(ctx context.Context, user models.User) error {
	_, err := m.client.UpdateByID(ctx, user.ID, user)
	if err != nil {
		return fmt.Errorf("failed to update user, %w", err)
	}
	return nil
}

// Delete deletes user entry.
func (m *Mongo) Delete(ctx context.Context, id int) error {
	_, err := m.client.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return fmt.Errorf("failed to delet user, %w", err)
	}
	return nil
}
