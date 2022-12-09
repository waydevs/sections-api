package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	designPatternsCollectionName = "design_patterns"
)

// DesignPatterns is a repository for DesignPattern.
type DesignPatterns struct {
	db DatabaseHelper
}

// NewDesignPatterns creates a new DesignPatterns repository.
func NewDesignPatterns(db DatabaseHelper) *DesignPatterns {
	return &DesignPatterns{db: db}
}

// GetByID returns a DesignPattern by its ID.
func (s *DesignPatterns) GetByID(ctx context.Context, id string) (DesignPattern, error) {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return DesignPattern{}, err
	}

	result := s.db.Collection(designPatternsCollectionName).FindOne(ctx, map[string]primitive.ObjectID{"_id": primitiveID})

	var designPattern DesignPattern
	err = result.Decode(&designPattern)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return DesignPattern{}, ErrNotFound
		}
		return DesignPattern{}, err
	}

	return designPattern, nil
}

// Create creates a new DesignPattern.
func (s *DesignPatterns) Create(ctx context.Context, designPattern DesignPattern) (DesignPattern, error) {
	result, err := s.db.Collection(designPatternsCollectionName).InsertOne(ctx, designPattern)
	if err != nil {
		return DesignPattern{}, err
	}

	designPattern.MongoID = result.(primitive.ObjectID)
	return designPattern, nil
}

// Delete deletes a DesignPattern by its ID.
func (d *DesignPatterns) Delete(ctx context.Context, id string) error {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = d.db.Collection(designPatternsCollectionName).DeleteOne(ctx, map[string]primitive.ObjectID{"_id": primitiveID})
	return err
}

// Update updates a DesignPattern.
func (d *DesignPatterns) Update(ctx context.Context, designPattern DesignPattern) (DesignPattern, error) {
	_, err := d.db.Collection(designPatternsCollectionName).ReplaceOne(ctx, map[string]primitive.ObjectID{"_id": designPattern.MongoID}, designPattern)
	if err != nil {
		return DesignPattern{}, err
	}

	return designPattern, nil
}
