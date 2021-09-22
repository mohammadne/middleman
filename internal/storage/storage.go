package storage

import (
	"github.com/mohammadne/middleman/internal/models"
)

type Storage interface {
	Save(hashId string, body *models.Body) error
	Get(hashId string) (*models.Body, error)
}
