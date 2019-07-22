package controllers

import (
	"finder/config"
	"finder/model"
	"finder/service"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
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
	s.Echo.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	//error handler
	s.Echo.HTTPErrorHandler = customHTTPErrorHandler

	// Routes
	/*
		e.POST("/users", controllers.CreateUser)
		e.GET("/users/:id", controllers.GetUser)
		e.PUT("/users/:id", controllers.UpdateUser)
		e.DELETE("/users/:id", controllers.DeleteUser)
	*/

	s.Echo.POST("/login", Login)
	s.Echo.POST("/logout", Logout)
	s.Echo.GET("/comment/:id/", GetCommentList)
	s.Echo.GET("/search", SearchRecordDetail)
	s.Echo.GET("/record/:id", GetRecordDetail)
	s.Echo.GET("/recordList", GetRecordList)

	admin := s.Echo.Group("/admin", MiddlewareAuthAdmin)
	admin.GET("/admin/userList", GetUserList)
	admin.POST("/admin/user", AddUser)
	admin.GET("/admin/user/:id", GetUser)
	admin.PUT("/admin/user/:id", UpdateUser)
	admin.DELETE("/admin/user/:id", DeleteUser)
	admin.PUT("/admin/user_password/:id", UpdatePassword)

	s.Echo.Logger.Fatal(s.Echo.Start(c.HttpAddress))
}

func MiddlewareAuthAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := GetSessionId(c)
		if userId <= 0 {
			return RenderLogin(c)
		}

		userInfo, err := findSrv.GetUser(userId)
		if err != nil || userInfo == nil {
			return RenderLogin(c)
		}

		return next(c)
	}
}

func GetSessionId(c echo.Context) (userId int64) {
	sess, err := session.Get("session", c)
	if err != nil {
		return
	}
	return sess.Values["id"].(int64)
}

func SetSessionId(c echo.Context, userId int64) (err error) {
	sess, err := session.Get("session", c)
	if err != nil {
		return
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
	}
	sess.Values["id"] = userId
	return sess.Save(c.Request(), c.Response())
}

func DelSessionId(c echo.Context) (err error) {
	sess, err := session.Get("session", c)
	if err != nil {
		return
	}

	sess.Options.MaxAge = -1

	return sess.Save(c.Request(), c.Response())
}

func GetUserRole(c echo.Context) (role int) {
	sess, err := session.Get("session", c)
	if err != nil {
		return
	}

	result, err := findSrv.GetUser(sess.Values["id"].(int64))
	if err != nil || result == nil {
		role = model.UserError
	}

	if result.Status == model.StatusERR {
		role = model.UserError
	}

	return
}

func RenderLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}

func RenderAdmin(c echo.Context) error {
	return c.Render(http.StatusOK, "admin", nil)
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
