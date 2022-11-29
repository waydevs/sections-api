package repository

import "fmt"

// DesignPattern handle calls to db related Design Patterns
type DesignPatterns struct {
	repo *Repository
}

// NewDesignPattern return a new instance of DesignPattern
func NewDesignPattern(repo *Repository) *DesignPatterns {
	return &DesignPatterns{
		repo: repo,
	}
}

// GetDesignPattern search a design pattern based on a key, and return it.
func (s *DesignPatterns) GetDesignPattern(key string) DesignPattern {
	response := s.repo.connExample.Get(key)

	return DesignPattern{
		Name: fmt.Sprintf("%s %s", response, "world"),
	}
}
