package storage

import (
	"github.com/mohammadne/middleman/internal/models"
	"github.com/mohammadne/middleman/pkg/logger"
)

type memoryStorage struct {
	logger    logger.Logger
	container map[string](*models.Body)
}

func NewMemoryStorage(logger logger.Logger) Storage {
	return &memoryStorage{logger: logger}
}

func (storage *memoryStorage) Save(hashId string, body *models.Body) error {
	return nil
}

func (storage *memoryStorage) Get(hashId string) (*models.Body, error) {
	return nil, nil
}
