package controllers

import (
	"finder/model"
	"finder/util"
	"github.com/labstack/echo"
	"html/template"
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

	result, err := findSrv.GetRecordList(page, pageSize, model.NotFind, model.ReviewSuccess)
	if err != nil {
		findSrv.Logger.Errorf("get record err:%s", err.Error())
		return JsonServerError(c)
	}

	return JsonOk(c, result)
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

	comment, err := findSrv.GetAllCommentList(recordId)
	if err != nil {
		findSrv.Logger.Errorf("get comment err:%s", err.Error())
		return JsonServerError(c)
	}

	captchaId, captchaImage := util.GenerateCaptcha()

	res := map[string]interface{}{
		"title":        "寻人信息",
		"data":         nil,
		"comment":      nil,
		"captchaId":    captchaId,
		"captchaImage": template.URL(captchaImage),
	}

	if result != nil {
		res["data"] = result
	}

	if comment != nil {
		for k, v := range comment {
			t := time.Unix(v.CreateTime, 0)
			comment[k].ShowTime = t.Format("2006-01-02 15:04:05")
		}
		res["comment"] = comment
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

	return c.Render(http.StatusOK, "add_record", res)
}

func RenderAdminRecord(c echo.Context) error {
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

	//if role == model.USER {
	if result.PublisherId != GetSessionId(c) {
		return JsonBadRequest(c, "权限不足")
	}
	//}

	res := map[string]interface{}{
		"userId":   GetSessionId(c),
		"name":     GetSessionName(c),
		"role":     GetUserRole(c),
		"title":    "修改寻人状态",
		"leftMenu": "record",
		"menu":     "user_record_list",
		"data":     nil,
	}

	if result != nil {
		res["data"] = result
	}

	return c.Render(http.StatusOK, "admin_record", res)
}

func RenderReviewRecord(c echo.Context) error {
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

	result, err := findSrv.GetRecord(recordId)
	if err != nil {
		findSrv.Logger.Errorf("get record err:%s", err.Error())
		return JsonServerError(c)
	}

	if result == nil {
		return JsonBadRequest(c, "id不正确")
	}

	res := map[string]interface{}{
		"userId":   GetSessionId(c),
		"name":     GetSessionName(c),
		"role":     GetUserRole(c),
		"title":    "审核寻人信息",
		"leftMenu": "record",
		"menu":     "record_list",
		"data":     nil,
	}

	if result != nil {
		res["data"] = result
	}

	return c.Render(http.StatusOK, "review_record", res)
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

	name := c.FormValue("name")
	sex := c.FormValue("sex")
	photo, err := c.FormFile("photo")
	if err != nil {
		return JsonBadRequest(c, "图片上传失败")
	}
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

	imgPath, err := util.UploadFile(photo)
	if err != nil {
		return JsonBadRequest(c, "图片上传失败")
	}

	record := &model.Record{
		PublisherId: GetSessionId(c),
		Name:        name,
		Sex:         tmpSex,
		Photo:       imgPath,
		Address:     address,
		Date:        date,
		Remark:      remark,
		IsFind:      model.NotFind,
		Status:      model.UnReview,
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

	//if role == model.USER {
	if result.PublisherId != GetSessionId(c) {
		return JsonBadRequest(c, "权限不足")
	}
	//}

	isFind := c.FormValue("isfind")

	if isFind == "" {
		return JsonBadRequest(c, "请选择发现状态")
	}

	tmpIsFind, _ := strconv.Atoi(isFind)

	record := &model.Record{
		IsFind:     tmpIsFind,
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

	count, err := findSrv.GetUserRecordCount(GetSessionId(c), isFind, status)
	if err != nil {
		findSrv.Logger.Errorf("GetUserRecordCount  err:%s", err.Error())
		return JsonServerError(c)
	}

	totalPage := int(math.Ceil(float64(count) / float64(pageSize))) //page总数
	if page > totalPage {
		page = totalPage
	}

	result, err := findSrv.GetUserRecordList(GetSessionId(c), page, pageSize, isFind, status)
	if err != nil {
		findSrv.Logger.Errorf("get record err:%s", err.Error())
		return JsonServerError(c)
	}

	res := map[string]interface{}{
		"userId":    GetSessionId(c),
		"name":      GetSessionName(c),
		"role":      GetUserRole(c),
		"title":     "我的寻人",
		"leftMenu":  "record",
		"menu":      "user_record_list",
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

	return c.Render(http.StatusOK, "user_record_list", res)
}
