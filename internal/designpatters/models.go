package designpatters

import "github.com/waydevs/sections-api/internal/platform/repository"

type DesignPattern struct {
	ID          string               `json:"id"`
	Title       string               `json:"title"`
	Subtitle    string               `json:"subtitle"`
	ContentData []repository.Content `json:"contentData"`
}
