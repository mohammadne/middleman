package storage

import "github.com/mohammadne/middleman/internal/models"

type Storage interface {
	Save(body models.Body)
	Get(hash string)
}

type storage struct{}

func New(directory string) Storage {
	return &storage{}
}

func (s *storage) Save(body models.Body) {}

func (s *storage) Get(hash string) {}
