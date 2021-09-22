package server

import (
	"log"

	"github.com/labstack/echo"
	"github.com/mohammadne/middleman/internal/handlers"
	"github.com/mohammadne/middleman/pkg/file"
)

// subDirectory := fmt.Sprintf("%s/%d", directory, port)
// addr := fmt.Sprintf("%s:%d", host, port)
// file.CreateDirIfMissed(subDirectory)
func Setup(directory string, path string) {
	err := file.CreateDirIfMissed(directory)
	if err != nil {
		log.Fatal("Error creating server directory")
	}

	e := echo.New()
	e.HideBanner = true
	handlers.SetupServerRoutes(directory, e)

	err = e.Start(path)
	if err != nil {
		log.Fatal(err)
	}
}
