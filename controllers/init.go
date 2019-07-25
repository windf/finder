package controllers

import (
	"finder/config"
	"finder/model"
	"finder/service"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
)

var findSrv *service.Service

type Message struct {
	Errcode int         `json:"errcode" xml:"errcode"`
	Errmsg  string      `json:"errmsg,omitempty" xml:"errmsg"`
	Data    interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Init(c *config.Config, s *service.Service) {
	findSrv = s

	//Middleware
	s.Echo.Use(middleware.Recover())
	s.Echo.Use(middleware.Logger())
	s.Echo.Use(middleware.RequestID())
	s.Echo.Use(session.Middleware(sessions.NewCookieStore([]byte(c.Secret))))

	//template
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	s.Echo.Renderer = t

	//error handler
	s.Echo.HTTPErrorHandler = customHTTPErrorHandler

	//静态目录
	s.Echo.Static("/static", "static")

	// Routes
	s.Echo.GET("/", RenderIndex)
	s.Echo.GET("/login", RenderLogin)
	s.Echo.POST("/login", Login)
	s.Echo.GET("/logout", Logout)
	s.Echo.POST("/logout", Logout)
	s.Echo.GET("/register", RenderRegister)
	s.Echo.POST("/register", Register)
	s.Echo.GET("/comment/:id", GetCommentList)
	s.Echo.POST("/comment/:id", AddComment)
	s.Echo.GET("/search", SearchRecordDetail)
	s.Echo.GET("/record/:id", RenderRecord)
	s.Echo.GET("/recordList", RenderRecordList)
	s.Echo.GET("/404", Render404)

	admin := s.Echo.Group("/admin", MiddlewareAuthAdmin)
	//index
	admin.GET("", RenderAdmin)
	//user
	admin.GET("/userList", RenderUserList)
	admin.GET("/user", RenderAddUser)
	admin.POST("/user", AddUser)
	admin.GET("/user/:id", RenderUser)
	admin.PUT("/user/:id", UpdateUser)
	admin.DELETE("/user/:id", DeleteUser)
	admin.GET("/user/:id/password", RenderPassword)
	admin.PUT("/user/:id/password", UpdatePassword)
	//record
	admin.POST("/record", AddRecord)
	admin.PUT("/record/:id", UpdateRecord)
	admin.DELETE("/record/:id", DeleteRecord)
	admin.PUT("/record/:id/review", ReviewRecord)
	admin.GET("/recordList", GetAdminRecordList)
	admin.GET("/user/recordList", GetUserRecordList)
	//comment
	admin.GET("/commentList", GetAdminCommentList)
	admin.DELETE("/comment/:id", DeleteComment)
	s.Echo.Logger.Fatal(s.Echo.Start(c.HttpAddress))
}

func MiddlewareAuthAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := GetSessionId(c)
		if userId <= 0 {
			return RedirectLogin(c)
		}

		userInfo, err := findSrv.GetUser(userId)
		if err != nil || userInfo == nil {
			return RedirectLogin(c)
		}

		return next(c)
	}
}

func GetSessionId(c echo.Context) (userId int64) {
	sess, err := session.Get("session", c)
	if err != nil {
		return
	}

	if id, ok := sess.Values["id"].(int64); ok {
		return id
	}

	return
}

func GetSessionName(c echo.Context) (name string) {
	sess, err := session.Get("session", c)
	if err != nil {
		return
	}

	if tmp, ok := sess.Values["name"].(string); ok {
		return tmp
	}

	return
}

func SetSessionId(c echo.Context, userId int64, name string) (err error) {
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
	sess.Values["name"] = name
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
	role = result.Role
	return
}

//重定向
func RedirectLogin(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/login")
}

func RedirectAdmin(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/admin")
}

func Redirect404(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/404")
}

func Render404(c echo.Context) error {
	return c.Render(http.StatusOK, "404", nil)
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

func JsonServerError(c echo.Context) error {
	code := http.StatusInternalServerError
	return c.JSON(code, &Message{
		Errcode: code,
		Errmsg:  "内部错误，请稍后重试",
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
			Redirect404(c)
		default:
			JsonServerError(c)
		}
	}
}
