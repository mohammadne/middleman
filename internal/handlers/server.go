package handlers

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/mohammadne/middleman/pkg/file"
	"github.com/mohammadne/middleman/pkg/model"
	"github.com/mohammadne/middleman/pkg/utils"
)

type ServerHandler struct {
	directory string
}

// SetupRoutes defines all routes needed on given `Echo` router
func SetupServerRoutes(directory string, e *echo.Echo) {
	handler := ServerHandler{directory}

	e.GET("/objects/:id", handler.Get)
	e.POST("/objects", handler.Post)
}

// GetObject will search associated directory and will
// return object if exists
func (sh ServerHandler) Get(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, parseErr := strconv.ParseInt(idStr, 10, 64)
	if parseErr != nil {
		code := http.StatusBadRequest
		return ctx.String(code, http.StatusText(code))
	}

	path := sh.directory + "/" + strconv.FormatInt(id, 10)

	if !file.IsFileExists(path) {
		code := http.StatusBadRequest
		return ctx.String(code, http.StatusText(code))
	}

	bytes, err := file.ReadFile(path)
	if err != nil {
		code := http.StatusBadRequest
		return ctx.String(code, err.Error())
	}

	values := strings.Split(string(bytes), "\n")

	return ctx.JSON(
		http.StatusOK,
		&model.Body{
			Key:   values[0],
			Value: values[1],
		},
	)

}

// PostObject will store body in file if exists before
func (sh ServerHandler) Post(ctx echo.Context) error {
	body := new(model.Body)
	if err := ctx.Bind(body); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	md5Key := md5.Sum([]byte(body.Key))
	md5Int := utils.ByteArrayToInt(md5Key)
	path := sh.directory + "/" + strconv.Itoa(md5Int)

	fmt.Println(path)
	if file.IsFileExists(path) {
		code := http.StatusBadRequest
		return ctx.String(code, http.StatusText(code))
	}

	f, err := file.CreateFile(path)
	if err != nil {
		code := http.StatusBadRequest
		return ctx.String(code, err.Error())
	}
	defer f.Close()
	f.WriteString(body.Key)
	f.WriteString("\n")
	f.WriteString(body.Value)

	return ctx.String(http.StatusCreated, "created")
}
