package designpatters

import (
	"fmt"

	"github.com/waydevs/sections-api/internal/platform/repository"
)

type DesignPattersService struct {
	repoDP *repository.DesignPatterns
}

func NewDesignPattersService(repository *repository.DesignPatterns) *DesignPattersService {
	return &DesignPattersService{
		repoDP: repository,
	}
}

func (s *DesignPattersService) GetPattern(key string) DesignPatternDTO {
	response := s.repoDP.GetDesignPattern(key)

	fmt.Println(response)

	return repositoryToDTO(response)
}

func repositoryToDTO(model repository.DesignPattern) DesignPatternDTO {
	return DesignPatternDTO{
		Name: model.Name,
	}
}
