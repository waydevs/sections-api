package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DesignPattern struct {
	MongoID     primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `json:"title"`
	Subtitle    string             `json:"subtitle"`
	ContentData []Content          `json:"contentData"`
}

type Content struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Image       []string `json:"image"`
}
