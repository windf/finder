package controllers

import (
	"finder/config"
	"finder/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

var findSrv *service.Service

type Message struct {
	Errcode int         `json:"errcode" xml:"errcode"`
	Errmsg  string      `json:"errmsg,omitempty" xml:"errmsg"`
	Data    interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func Init(c *config.Config, s *service.Service) {
	findSrv = s

	//Middleware
	s.Echo.Use(middleware.Recover())
	s.Echo.Use(middleware.Logger())
	s.Echo.Use(middleware.RequestID())

	//error handler
	s.Echo.HTTPErrorHandler = customHTTPErrorHandler

	// Routes
	/*
		e.POST("/users", controllers.CreateUser)
		e.GET("/users/:id", controllers.GetUser)
		e.PUT("/users/:id", controllers.UpdateUser)
		e.DELETE("/users/:id", controllers.DeleteUser)
	*/
	s.Echo.GET("/record/:id", GetRecordDetail)
	s.Echo.GET("/recordList/", GetRecordList)

	s.Echo.Logger.Fatal(s.Echo.Start(c.HttpAddress))
}

func JsonOk(c echo.Context, i interface{}) error {
	code := http.StatusOK
	return c.JSON(code, &Message{
		Errcode: code,
		Errmsg:  "success",
		Data:    i,
	})
}

func JsonBadRequest(c echo.Context, msg string) error {
	code := http.StatusBadRequest
	return c.JSON(code, &Message{
		Errcode: code,
		Errmsg:  msg,
	})
}

func JsonError(c echo.Context, code int, msg string) error {
	return c.JSON(code, &Message{
		Errcode: code,
		Errmsg:  msg,
	})
}

func customHTTPErrorHandler(err error, c echo.Context) {
	httpError, ok := err.(*echo.HTTPError)
	if ok {
		errorCode := httpError.Code
		switch errorCode {
		case http.StatusNotFound:
			JsonError(c, http.StatusNotFound, "Not Found")
		default:
			JsonError(c, http.StatusInternalServerError, "内部错误")
		}
	}
}
