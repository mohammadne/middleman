package storage

import (
	"github.com/mohammadne/middleman/internal/models"
)

type Storage interface {
	Save(filename string, body *models.Body) error
	Get(hash string) (*models.Body, error)
}
