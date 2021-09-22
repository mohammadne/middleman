package handlers

import (
	"crypto/md5"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/mohammadne/middleman/pkg/model"
	"github.com/mohammadne/middleman/pkg/request"
	"github.com/mohammadne/middleman/pkg/utils"
)

type LoadBalancerHandler struct {
	servers []string
	cache   map[string](*model.Body)
}

// SetupRoutes defines all routes needed on given `Echo` router
func SetupLoadBalancerRoutes(servers []string, e *echo.Echo) {
	handler := LoadBalancerHandler{
		servers: servers,
		cache:   map[string]*model.Body{},
	}

	e.GET("/objects/:id", handler.Get)
	e.POST("/objects", handler.Post)
}

// GetObject will search associated directory and will
// return object if exists
func (lbh LoadBalancerHandler) Get(ctx echo.Context) error {
	strId := ctx.Param("id")
	_, parseErr := strconv.ParseInt(strId, 10, 64)
	if parseErr != nil {
		code := http.StatusBadRequest
		return ctx.String(code, http.StatusText(code))
	}

	value, ok := lbh.cache[strId]
	if !ok {
		md5Int, _ := strconv.Atoi(strId)
		remain := md5Int % len(lbh.servers)

		return request.Get(lbh.servers[remain], nil)
	}

	return ctx.JSON(http.StatusOK, &value)
}

// PostObject will store body in file if exists before
func (lbh LoadBalancerHandler) Post(ctx echo.Context) error {
	body := new(model.Body)
	if err := ctx.Bind(body); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	md5Key := md5.Sum([]byte(body.Key))
	md5Str := string(md5Key[:])

	if _, ok := lbh.cache[md5Str]; body.Cache && !ok {
		lbh.cache[md5Str] = body
	}

	md5Int := utils.ByteArrayToInt(md5Key)
	remain := md5Int % len(lbh.servers)

	return request.Post("http://"+lbh.servers[remain]+"/objects", &body)
}
