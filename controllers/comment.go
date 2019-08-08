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

func GetCommentList(c echo.Context) error {
	id := c.Param("id")
	reqPage := c.QueryParam("page")
	reqPageSize := c.QueryParam("pageSize")

	recordId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	if recordId <= 0 {
		return JsonBadRequest(c, "id不正确")
	}

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

	count, err := findSrv.GetCommentCount(recordId)
	if err != nil {
		findSrv.Logger.Errorf("GetCommentCount  err:%s", err.Error())
		return JsonServerError(c)
	}

	totalPage := int(math.Ceil(float64(count) / float64(pageSize))) //page总数
	if page > totalPage {
		page = totalPage
	}

	result, err := findSrv.GetCommentList(recordId, page, pageSize)
	if err != nil {
		findSrv.Logger.Errorf("GetCommentList err:%s", err.Error())
		return JsonServerError(c)
	}

	var res interface{}
	res = struct{}{}

	if result != nil {
		res = result
	}

	return JsonOk(c, res)
}

func AddComment(c echo.Context) error {
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

	if result == nil {
		return JsonBadRequest(c, "id不正确")
	}

	//无需注册提供线索
	phone := c.FormValue("phone")
	remark := c.FormValue("remark")
	photo, err := c.FormFile("photo")
	if err != nil {
		return JsonBadRequest(c, "图片上传失败")
	}

	if phone == "" {
		return JsonBadRequest(c, "请输入联系方式")
	}

	if remark == "" {
		return JsonBadRequest(c, "请输入线索描述")
	}

	imgPath, err := util.UploadFile(photo)
	if err != nil {
		return JsonBadRequest(c, "图片上传失败")
	}

	comment := &model.Comment{
		RecordId:   recordId,
		Phone:      phone,
		Photo:      imgPath,
		Remark:     remark,
		CreateTime: time.Now().Unix(),
	}

	if !findSrv.AddComment(comment) {
		return JsonBadRequest(c, "添加评论失败，请稍后重试")
	}

	return JsonOk(c, struct{}{})
}

func GetAdminCommentList(c echo.Context) error {
	role := GetUserRole(c)
	if role <= model.USER {
		return JsonBadRequest(c, "权限不足")
	}

	reqPage := c.QueryParam("page")
	reqPageSize := c.QueryParam("pageSize")
	reqId := c.QueryParam("record_id")

	page, err := strconv.Atoi(reqPage)
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(reqPageSize)
	if err != nil {
		pageSize = 10
	}

	recordId, err := strconv.ParseInt(reqId, 10, 64)
	if err != nil {
		recordId = 0
	}

	if page < 1 {
		page = 1
	}

	if pageSize <= 0 || pageSize > 200 {
		pageSize = 10
	}

	if recordId < 1 {
		recordId = 0
	}

	count, err := findSrv.GetCommentCount(recordId)
	if err != nil {
		findSrv.Logger.Errorf("GetCommentCount  err:%s", err.Error())
		return JsonServerError(c)
	}

	totalPage := int(math.Ceil(float64(count) / float64(pageSize))) //page总数
	if page > totalPage {
		page = totalPage
	}

	result, err := findSrv.GetCommentList(recordId, page, pageSize)
	if err != nil {
		findSrv.Logger.Errorf("get GetCommentList err:%s", err.Error())
		return JsonServerError(c)
	}

	res := map[string]interface{}{
		"userId":    GetSessionId(c),
		"name":      GetSessionName(c),
		"role":      GetUserRole(c),
		"title":     "评论列表",
		"leftMenu":  "comment",
		"menu":      "comment_list",
		"data":      nil,
		"totalPage": totalPage,
		"pageSize":  pageSize,
		"page":      page,
		"recordId":  reqId,
	}

	if result != nil {
		res["data"] = result
	}

	return c.Render(http.StatusOK, "comment_list", res)
}

func DeleteComment(c echo.Context) (err error) {
	id := c.Param("id")
	commentId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	if commentId <= 0 {
		return JsonBadRequest(c, "id不正确")
	}

	role := GetUserRole(c)
	if role <= model.USER {
		return JsonBadRequest(c, "权限不足")
	}

	result, err := findSrv.GetComment(commentId)
	if err != nil {
		findSrv.Logger.Errorf("get GetComment err:%s", err.Error())
		return JsonServerError(c)
	}

	if result == nil {
		return JsonBadRequest(c, "id不正确")
	}

	if !findSrv.DeleteComment(&model.Comment{ID: commentId}) {
		return JsonBadRequest(c, "删除评论失败，请稍后重试")
	}

	return JsonOk(c, struct{}{})
}
