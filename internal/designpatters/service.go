package designpatters

import (
	"context"
	"errors"
	"fmt"

	"github.com/waydevs/sections-api/internal/platform/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// ErrSomethingWentWrong is returned when something went wrong.
	ErrSomethingWentWrong = errors.New("Something went wrong")

	// ErrDesignPatternNotFound is returned when a DesignPattern is not found.
	ErrDesignPatternNotFound = errors.New("Design Pattern not found")
)

// DesignPatternRepository is a repository for DesignPattern.
type DesignPatternRepository interface {
	GetByID(ctx context.Context, id string) (repository.DesignPattern, error)
	Create(ctx context.Context, designPattern repository.DesignPattern) (repository.DesignPattern, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, designPattern repository.DesignPattern) (repository.DesignPattern, error)
}

// Service handles the business logic and use cases for DesignPattern.
type Service struct {
	db DesignPatternRepository
}

// NewService creates a new DesignPattern service.
func NewService(db DesignPatternRepository) *Service {
	return &Service{db: db}
}

// GetByID returns a DesignPattern by its ID.
func (s *Service) GetByID(ctx context.Context, id string) (DesignPattern, error) {
	// Validaciones o cache

	designPattern, err := s.db.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return DesignPattern{}, ErrDesignPatternNotFound
		}

		fmt.Println(err)
		return DesignPattern{}, ErrSomethingWentWrong
	}

	return repositoryModelToServiceModel(designPattern), nil
}

// Create creates a new DesignPattern.
func (s *Service) Create(ctx context.Context, designPattern DesignPattern) (DesignPattern, error) {
	// Validaciones o cache

	convertedDesignPattern, err := serviceModelToRepositoryModelForCreation(designPattern)
	if err != nil {
		fmt.Println(err)
		return DesignPattern{}, ErrSomethingWentWrong
	}
	designPatternCreated, err := s.db.Create(ctx, convertedDesignPattern)
	if err != nil {
		fmt.Println(err)
		return DesignPattern{}, ErrSomethingWentWrong
	}

	return repositoryModelToServiceModel(designPatternCreated), nil
}

// Delete deletes a DesignPattern by its ID.
func (s *Service) Delete(ctx context.Context, id string) error {
	// Validaciones o cache

	err := s.db.Delete(ctx, id)
	if err != nil {
		fmt.Println(err)
		return ErrSomethingWentWrong
	}

	return nil
}

// Update updates a DesignPattern.
func (s *Service) Update(ctx context.Context, designPattern DesignPattern) (DesignPattern, error) {
	// Validaciones o cache

	convertedDesignPattern, err := serviceModelToRepositoryModel(designPattern)
	if err != nil {
		fmt.Println(err)
		return DesignPattern{}, ErrSomethingWentWrong
	}

	designPatternUpdated, err := s.db.Update(ctx, convertedDesignPattern)
	if err != nil {
		fmt.Println(err)
		return DesignPattern{}, ErrSomethingWentWrong
	}

	return repositoryModelToServiceModel(designPatternUpdated), nil
}

// Hacemos la converson de los modelos de la capa de repositorio a los modelos de la capa de servicio
// y viceversa porque no queremos que la capa de servicio tenga que depender de la capa de repositorio
// ni devolver modelos de la capa de repositorio al usuario.

func repositoryModelToServiceModel(designPattern repository.DesignPattern) DesignPattern {
	return DesignPattern{
		ID:          designPattern.MongoID.Hex(),
		Title:       designPattern.Title,
		Subtitle:    designPattern.Subtitle,
		ContentData: designPattern.ContentData,
	}
}

func serviceModelToRepositoryModelForCreation(designPattern DesignPattern) (repository.DesignPattern, error) {
	return repository.DesignPattern{
		Title:       designPattern.Title,
		Subtitle:    designPattern.Subtitle,
		ContentData: designPattern.ContentData,
	}, nil
}

func serviceModelToRepositoryModel(designPattern DesignPattern) (repository.DesignPattern, error) {
	primitiveID, err := primitive.ObjectIDFromHex(designPattern.ID)
	if err != nil {
		return repository.DesignPattern{}, err
	}

	return repository.DesignPattern{
		MongoID:     primitiveID,
		Title:       designPattern.Title,
		Subtitle:    designPattern.Subtitle,
		ContentData: designPattern.ContentData,
	}, nil
}
