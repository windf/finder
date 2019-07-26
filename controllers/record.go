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

func RenderRecordList(c echo.Context) error {
	reqPage := c.QueryParam("page")
	reqPageSize := c.QueryParam("pageSize")

	page, err := strconv.Atoi(reqPage)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	pageSize, err := strconv.Atoi(reqPageSize)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	if page < 1 {
		page = 1
	}

	if pageSize <= 0 || pageSize > 200 {
		pageSize = 10
	}

	result, err := findSrv.GetRecordList(page, pageSize, model.AllFind, model.ReviewSuccess)
	if err != nil {
		findSrv.Logger.Errorf("get record err:%s", err.Error())
		return JsonServerError(c)
	}

	var res interface{}
	res = struct{}{}

	if result != nil {
		res = result
	}

	return c.Render(http.StatusOK, "record_list", res)
	//return JsonOk(c, res)
}

func RenderRecord(c echo.Context) (err error) {
	id := c.Param("id")
	recordId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	if recordId <= 0 {
		return JsonBadRequest(c, "id不正确")
	}

	result, err := findSrv.GetRecord(recordId)
	if err != nil {
		findSrv.Logger.Errorf("get record err:%s", err.Error())
		return JsonServerError(c)
	}

	var res interface{}
	res = struct{}{}

	if result != nil {
		res = result
	}
	return c.Render(http.StatusOK, "record", res)
}

func RenderAddRecord(c echo.Context) error {
	res := map[string]interface{}{
		"userId":   GetSessionId(c),
		"name":     GetSessionName(c),
		"role":     GetUserRole(c),
		"title":    "添加寻人",
		"leftMenu": "record",
		"menu":     "add_record",
	}

	return c.Render(http.StatusOK, "record", res)
}

func RenderAdminRecord(c echo.Context) error {
	res := map[string]interface{}{
		"userId":   GetSessionId(c),
		"name":     GetSessionName(c),
		"role":     GetUserRole(c),
		"title":    "修改寻人信息",
		"leftMenu": "record",
		"menu":     "admin_record",
	}

	return c.Render(http.StatusOK, "admin_record", res)
}

func SearchRecordDetail(c echo.Context) (err error) {
	keyword := c.QueryParam("keyword")

	id, ok := util.IsNumeric(keyword)

	var res interface{}
	res = struct{}{}

	if ok {
		if id <= 0 {
			return JsonBadRequest(c, "id不正确")
		}
		result, err := findSrv.GetRecord(id)
		if err != nil {
			findSrv.Logger.Errorf("get record err:%s", err.Error())
			return JsonServerError(c)
		}

		if result != nil {
			res = []*model.Record{result}
		}
	} else {
		if keyword == "" {
			return JsonBadRequest(c, "参数错误")
		}

		result, err := findSrv.GetRecordByName(keyword)
		if err != nil {
			findSrv.Logger.Errorf("GetRecordByName err:%s", err.Error())
			return JsonServerError(c)
		}

		if result != nil {
			res = result
		}
	}

	return JsonOk(c, res)
}

func AddRecord(c echo.Context) (err error) {
	role := GetUserRole(c)
	if role == model.UserError {
		return JsonBadRequest(c, "权限不足")
	}

	//todo  上传图片
	name := c.FormValue("name")
	sex := c.FormValue("sex")
	photo := c.FormValue("photo")
	address := c.FormValue("address")
	date := c.FormValue("date")
	remark := c.FormValue("remark")

	if name == "" {
		return JsonBadRequest(c, "请输入姓名")
	}

	if sex == "" {
		return JsonBadRequest(c, "请输入性别")
	}

	if address == "" {
		return JsonBadRequest(c, "请输入地址")
	}

	if date == "" {
		return JsonBadRequest(c, "请输入日期")
	}

	tmpSex, _ := strconv.Atoi(sex)

	record := &model.Record{
		PublisherId: GetSessionId(c),
		Name:        name,
		Sex:         tmpSex,
		Photo:       photo,
		Address:     address,
		Date:        date,
		Remark:      remark,
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now().Unix(),
	}

	if !findSrv.AddRecord(record) {
		return JsonBadRequest(c, "添加信息失败，请稍后重试")
	}

	return JsonOk(c, struct{}{})
}

func UpdateRecord(c echo.Context) (err error) {
	id := c.Param("id")
	recordId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	if recordId <= 0 {
		return JsonBadRequest(c, "id不正确")
	}

	role := GetUserRole(c)
	if role == model.UserError {
		return JsonBadRequest(c, "权限不足")
	}

	result, err := findSrv.GetRecord(recordId)
	if err != nil {
		findSrv.Logger.Errorf("get record err:%s", err.Error())
		return JsonServerError(c)
	}

	if result == nil {
		return JsonBadRequest(c, "id不正确")
	}

	if role == model.USER {
		if result.PublisherId != GetSessionId(c) {
			return JsonBadRequest(c, "权限不足")
		}
	}

	name := c.FormValue("name")
	sex := c.FormValue("sex")
	photo := c.FormValue("photo")
	address := c.FormValue("address")
	date := c.FormValue("date")
	remark := c.FormValue("remark")

	if name == "" {
		return JsonBadRequest(c, "请输入姓名")
	}

	if sex == "" {
		return JsonBadRequest(c, "请输入性别")
	}

	if address == "" {
		return JsonBadRequest(c, "请输入地址")
	}

	if date == "" {
		return JsonBadRequest(c, "请输入日期")
	}

	tmpSex, _ := strconv.Atoi(sex)

	record := &model.Record{
		Name:       name,
		Sex:        tmpSex,
		Photo:      photo,
		Address:    address,
		Date:       date,
		Remark:     remark,
		UpdateTime: time.Now().Unix(),
	}

	if !findSrv.UpdateRecord(recordId, record) {
		return JsonBadRequest(c, "修改信息失败，请稍后重试")
	}

	return JsonOk(c, struct{}{})
}

func ReviewRecord(c echo.Context) (err error) {
	id := c.Param("id")
	recordId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	if recordId <= 0 {
		return JsonBadRequest(c, "id不正确")
	}

	role := GetUserRole(c)
	if role <= model.USER {
		return JsonBadRequest(c, "权限不足")
	}

	status := c.FormValue("status")
	if status == "" {
		return JsonBadRequest(c, "请选择审核状态")
	}

	tmpStatus, _ := strconv.Atoi(status)

	record := &model.Record{
		Status:     tmpStatus,
		UpdateTime: time.Now().Unix(),
	}

	if !findSrv.UpdateRecord(recordId, record) {
		return JsonBadRequest(c, "审核信息失败，请稍后重试")
	}

	return JsonOk(c, struct{}{})
}

func DeleteRecord(c echo.Context) (err error) {
	id := c.Param("id")
	recordId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	if recordId <= 0 {
		return JsonBadRequest(c, "id不正确")
	}

	role := GetUserRole(c)
	if role == model.UserError {
		return JsonBadRequest(c, "权限不足")
	}

	result, err := findSrv.GetRecord(recordId)
	if err != nil {
		findSrv.Logger.Errorf("get record err:%s", err.Error())
		return JsonServerError(c)
	}

	if result == nil {
		return JsonBadRequest(c, "id不正确")
	}

	if role == model.USER {
		if result.PublisherId != GetSessionId(c) {
			return JsonBadRequest(c, "权限不足")
		}
	}

	if !findSrv.DeleteRecord(&model.Record{ID: recordId}) {
		return JsonBadRequest(c, "删除信息失败，请稍后重试")
	}

	return JsonOk(c, struct{}{})
}

func RenderAdminRecordList(c echo.Context) error {
	role := GetUserRole(c)
	if role <= model.USER {
		return JsonBadRequest(c, "权限不足")
	}

	reqPage := c.QueryParam("page")
	reqPageSize := c.QueryParam("pageSize")
	reqIsFind := c.QueryParam("isfind")
	reqStatus := c.QueryParam("status")

	page, err := strconv.Atoi(reqPage)
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(reqPageSize)
	if err != nil {
		pageSize = 10
	}

	isFind, err := strconv.Atoi(reqIsFind)
	if err != nil {
		isFind = model.AllFind
	}

	status, err := strconv.Atoi(reqStatus)
	if err != nil {
		status = model.AllReview
	}

	if page < 1 {
		page = 1
	}

	if pageSize <= 0 || pageSize > 200 {
		pageSize = 10
	}

	if isFind < model.AllFind || isFind > model.FindOK {
		isFind = model.AllFind
	}

	if status < model.AllReview || status > model.ReviewSuccess {
		status = model.AllReview
	}

	count, err := findSrv.GetRecordCount(isFind, status)
	if err != nil {
		findSrv.Logger.Errorf("GetRecordCount  err:%s", err.Error())
		return JsonServerError(c)
	}

	totalPage := int(math.Ceil(float64(count) / float64(pageSize))) //page总数
	if page > totalPage {
		page = totalPage
	}

	result, err := findSrv.GetRecordList(page, pageSize, isFind, status)
	if err != nil {
		findSrv.Logger.Errorf("get record err:%s", err.Error())
		return JsonServerError(c)
	}

	res := map[string]interface{}{
		"userId":    GetSessionId(c),
		"name":      GetSessionName(c),
		"role":      GetUserRole(c),
		"title":     "寻人列表",
		"leftMenu":  "record",
		"menu":      "record_list",
		"data":      nil,
		"totalPage": totalPage,
		"pageSize":  pageSize,
		"page":      page,
		"isFind":    isFind,
		"status":    status,
	}

	if result != nil {
		res["data"] = result
	}

	return c.Render(http.StatusOK, "record_list", res)
}

func GetUserRecordList(c echo.Context) error {
	reqPage := c.QueryParam("page")
	reqPageSize := c.QueryParam("pageSize")
	reqIsFind := c.QueryParam("isfind")

	page, err := strconv.Atoi(reqPage)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	pageSize, err := strconv.Atoi(reqPageSize)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	isFind, err := strconv.Atoi(reqIsFind)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	if page < 1 {
		page = 1
	}

	if pageSize <= 0 || pageSize > 200 {
		pageSize = 10
	}

	if isFind < model.AllFind || isFind > model.FindOK {
		isFind = model.AllFind
	}

	result, err := findSrv.GetUserRecordList(GetSessionId(c), page, pageSize, isFind)
	if err != nil {
		findSrv.Logger.Errorf("get record err:%s", err.Error())
		return JsonServerError(c)
	}

	var res interface{}
	res = struct{}{}

	if result != nil {
		res = result
	}

	return JsonOk(c, res)
}
