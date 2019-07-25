package controllers

import (
	"finder/model"
	"finder/util"
	"github.com/labstack/echo"
	"math"
	"net/http"
	"strconv"
	"time"
)

func RenderUserList(c echo.Context) error {
	role := GetUserRole(c)
	if role <= model.USER {
		return JsonBadRequest(c, "权限不足")
	}

	reqPage := c.QueryParam("page")
	reqPageSize := c.QueryParam("pageSize")

	page, err := strconv.Atoi(reqPage)
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(reqPageSize)
	if err != nil {
		pageSize = 10
	}

	if page < 1 {
		page = 1
	}

	if pageSize <= 0 || pageSize > 200 {
		pageSize = 10
	}

	count, err := findSrv.GetUserCount()
	if err != nil {
		findSrv.Logger.Errorf("GetUserCount  err:%s", err.Error())
		return JsonServerError(c)
	}

	totalPage := int(math.Ceil(float64(count) / float64(pageSize))) //page总数
	if page > totalPage {
		page = totalPage
	}

	result, err := findSrv.GetUserList(page, pageSize)
	if err != nil {
		findSrv.Logger.Errorf("get user list err:%s", err.Error())
		return JsonServerError(c)
	}

	res := map[string]interface{}{
		"userId":    GetSessionId(c),
		"name":      GetSessionName(c),
		"role":      GetUserRole(c),
		"title":     "用户列表",
		"leftMenu":  "user",
		"menu":      "user_list",
		"data":      nil,
		"totalPage": totalPage,
		"pageSize":  pageSize,
		"page":      page,
	}

	if result != nil {
		res["data"] = result
	}

	return c.Render(http.StatusOK, "user_list", res)
}

func RenderAddUser(c echo.Context) error {
	return c.Render(http.StatusOK, "add_user", nil)
}

func RenderPassword(c echo.Context) error {
	return c.Render(http.StatusOK, "password", nil)
}

func RenderUser(c echo.Context) error {
	id := c.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	if userId <= 0 {
		return JsonBadRequest(c, "id不正确")
	}

	role := GetUserRole(c)
	if role <= model.USER {
		tmpUserId := GetSessionId(c)
		if userId != tmpUserId {
			return JsonBadRequest(c, "权限不足")
		}
	}

	result, err := findSrv.GetUser(userId)
	if err != nil {
		findSrv.Logger.Errorf("get user err:%s", err.Error())
		return JsonServerError(c)
	}

	res := map[string]interface{}{
		"userId":   GetSessionId(c),
		"name":     GetSessionName(c),
		"role":     GetUserRole(c),
		"title":    "用户信息修改",
		"leftMenu": "user",
		"menu":     "user_list",
		"data":     nil,
		"isSelf":   0, //是自己可以修改信息 0否 1是
	}

	if role == model.USER {
		res["menu"] = "user_edit"
	}

	if res["userId"] == result.ID {
		res["isSelf"] = 1
	}

	if result != nil {
		res["data"] = result
	}

	return c.Render(http.StatusOK, "user", res)
}

func AddUser(c echo.Context) error {
	role := GetUserRole(c)
	if role != model.ADMIN {
		return JsonBadRequest(c, "权限不足")
	}

	userName := c.FormValue("user_name")
	password := c.FormValue("password")
	nickName := c.FormValue("nickname")
	phone := c.FormValue("phone")
	email := c.FormValue("email")
	remark := c.FormValue("remark")
	status := c.FormValue("status")
	tmpRole := c.FormValue("role")

	if userName == "" {
		return JsonBadRequest(c, "请输入用户名")
	}

	if password == "" {
		return JsonBadRequest(c, "请输入密码")
	}

	if findSrv.CheckUser(userName) {
		return JsonBadRequest(c, "用户名已经存在")
	}

	tmpS, _ := strconv.Atoi(status)
	tmpR, _ := strconv.Atoi(tmpRole)

	user := &model.User{
		UserName:   userName,
		Password:   util.Md5(password),
		NickName:   nickName,
		Phone:      phone,
		Email:      email,
		Remark:     remark,
		Status:     tmpS,
		Role:       tmpR,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	if !findSrv.AddUser(user) {
		return JsonBadRequest(c, "添加用户失败，请稍后重试")
	}

	return JsonOk(c, struct{}{})
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	if userId <= 0 {
		return JsonBadRequest(c, "id不正确")
	}

	role := GetUserRole(c)
	if role <= model.REVIEWER {
		tmpUserId := GetSessionId(c)
		if userId != tmpUserId {
			return JsonBadRequest(c, "权限不足")
		}
	}

	nickName := c.FormValue("nickname")
	phone := c.FormValue("phone")
	email := c.FormValue("email")
	remark := c.FormValue("remark")

	user := &model.User{
		NickName:   nickName,
		Phone:      phone,
		Email:      email,
		Remark:     remark,
		UpdateTime: time.Now().Unix(),
	}

	if role == model.ADMIN {
		status := c.FormValue("status")
		tmpRole := c.FormValue("role")
		tmpS, _ := strconv.Atoi(status)
		tmpR, _ := strconv.Atoi(tmpRole)
		user.Status = tmpS
		user.Role = tmpR
	}

	if !findSrv.UpdateUser(userId, user) {
		return JsonBadRequest(c, "修改用户失败，请稍后重试")
	}
	return JsonOk(c, struct{}{})
}

func UpdatePassword(c echo.Context) error {
	id := c.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	if userId <= 0 {
		return JsonBadRequest(c, "id不正确")
	}

	role := GetUserRole(c)
	if role <= model.REVIEWER {
		tmpUserId := GetSessionId(c)
		if userId != tmpUserId {
			return JsonBadRequest(c, "权限不足")
		}
	}

	password := c.FormValue("password")
	newPassword := c.FormValue("new_password")
	rePassword := c.FormValue("re_password")

	if password == "" {
		return JsonBadRequest(c, "请输入原密码")
	}

	result, _ := findSrv.GetUser(userId)
	if result.Password != util.Md5(password) {
		return JsonBadRequest(c, "原密码不正确")
	}

	if newPassword == "" {
		return JsonBadRequest(c, "请输入新密码")
	}

	if newPassword != rePassword {
		return JsonBadRequest(c, "新密码和确认密码不一致")
	}

	user := &model.User{
		Password:   util.Md5(newPassword),
		UpdateTime: time.Now().Unix(),
	}

	if !findSrv.UpdateUser(userId, user) {
		return JsonBadRequest(c, "修改密码失败，请稍后重试")
	}

	return JsonOk(c, struct{}{})
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	if userId <= 0 {
		return JsonBadRequest(c, "id不正确")
	}

	role := GetUserRole(c)
	if role != model.ADMIN {
		return JsonBadRequest(c, "权限不足")
	}

	if userId == GetSessionId(c) {
		return JsonBadRequest(c, "不能删除自己")
	}

	if !findSrv.DeleteUser(&model.User{ID: userId}) {
		return JsonBadRequest(c, "删除用户失败，请稍后重试")
	}
	return JsonOk(c, struct{}{})
}
