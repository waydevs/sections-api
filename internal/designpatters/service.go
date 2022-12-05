package designpatters

import (
	"context"
	"fmt"

	"github.com/waydevs/sections-api/internal/platform/repository"
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
		fmt.Println(err)
		return DesignPattern{}, err
	}

	return repositoryModelToServiceModel(designPattern), nil
}

// Create creates a new DesignPattern.
func (s *Service) Create(ctx context.Context, designPattern DesignPattern) (DesignPattern, error) {
	// Validaciones o cache

	designPatternCreated, err := s.db.Create(ctx, serviceModelToRepositoryModel(designPattern))
	if err != nil {
		fmt.Println(err)
		return DesignPattern{}, err
	}

	return repositoryModelToServiceModel(designPatternCreated), nil
}

// Delete deletes a DesignPattern by its ID.
func (s *Service) Delete(ctx context.Context, id string) error {
	// Validaciones o cache

	return s.db.Delete(ctx, id)
}

// Update updates a DesignPattern.
func (s *Service) Update(ctx context.Context, designPattern DesignPattern) (DesignPattern, error) {
	// Validaciones o cache

	designPatternUpdated, err := s.db.Update(ctx, serviceModelToRepositoryModel(designPattern))
	if err != nil {
		fmt.Println(err)
		return DesignPattern{}, err
	}

	return repositoryModelToServiceModel(designPatternUpdated), nil
}

// Hacemos la converson de los modelos de la capa de repositorio a los modelos de la capa de servicio
// y viceversa porque no queremos que la capa de servicio tenga que depender de la capa de repositorio
// ni devolver modelos de la capa de repositorio al usuario.

func repositoryModelToServiceModel(designPattern repository.DesignPattern) DesignPattern {
	return DesignPattern{
		Title: designPattern.Title,
	}
}

func serviceModelToRepositoryModel(designPattern DesignPattern) repository.DesignPattern {
	return repository.DesignPattern{
		Title: designPattern.Title,
	}
}
