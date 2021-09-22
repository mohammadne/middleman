package storage

import (
	"fmt"

	"github.com/mohammadne/middleman/internal/models"
	"github.com/mohammadne/middleman/pkg/logger"
)

type memoryStorage struct {
	logger    logger.Logger
	container map[string](*models.Body)
}

func NewMemoryStorage(logger logger.Logger) Storage {
	container := make(map[string](*models.Body))
	return &memoryStorage{logger: logger, container: container}
}

func (storage *memoryStorage) Save(hashId string, body *models.Body) error {
	storage.container[hashId] = body
	return nil
}

func (storage *memoryStorage) Get(hashId string) (*models.Body, error) {
	value, ok := storage.container[hashId]
	if !ok {
		return nil, fmt.Errorf("value of : %s not found", hashId)
	}

	return value, nil
}
