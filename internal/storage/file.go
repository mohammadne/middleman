package storage

import (
	"fmt"
	"strings"

	"github.com/mohammadne/middleman/internal/models"
	"github.com/mohammadne/middleman/pkg/logger"
	"github.com/mohammadne/middleman/pkg/utils"
)

type fileStorage struct {
	logger    logger.Logger
	directory string
}

func NewFileStorage(directory string, logger logger.Logger) (Storage, error) {
	storage := &fileStorage{logger: logger}

	err := utils.CreateDirIfMissed(directory)
	if err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *fileStorage) Save(filename string, body *models.Body) error {
	path := s.directory + "/" + filename

	if utils.IsFileExists(path) {
		return fmt.Errorf("file: %s is already exists", path)
	}

	file, err := utils.CreateFile(path)
	if err != nil {
		return err
	}
	defer file.Close()

	content := fmt.Sprintf("%s\n%s", body.Key, body.Value)
	file.WriteString(content)

	return nil
}

func (s *fileStorage) Get(hashId string) (*models.Body, error) {
	// id, parseErr := strconv.ParseInt(idStr, 10, 64)
	// if parseErr != nil {
	// 	code := http.StatusBadRequest
	// 	return ctx.String(code, http.StatusText(code))
	// }

	// path := sh.directory + "/" + strconv.FormatInt(id, 10)

	path := fmt.Sprintf("%s/%s", s.directory, hashId)

	if !utils.IsFileExists(path) {
		return nil, fmt.Errorf("file: %s not exists", hashId)
	}

	bytes, err := utils.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bytes), "\n")
	return &models.Body{Key: lines[0], Value: lines[1]}, nil
}
