package file

import (
	"io/ioutil"
	"os"
)

func CreateDirIfMissed(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		return os.Mkdir(dirName, os.ModePerm)
	}

	return nil
}

func IsFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	if info.IsDir() {
		return false
	}

	return true
}

func ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func CreateFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	err = os.Chmod(filename, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return file, nil
}
