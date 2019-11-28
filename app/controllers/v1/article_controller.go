package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/varandrew/gin-product/app/models"
	"github.com/varandrew/gin-product/app/pkg/errno"
	"github.com/varandrew/gin-product/app/pkg/logging"
	"github.com/varandrew/gin-product/app/pkg/setting"
	"github.com/varandrew/gin-product/app/utils"
)

// @Summary 获取单个文章
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500  {string} string "{"code":500,"msg":"fail"}"
// @Router /api/v1/articles/{id} [get]
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := errno.INVALID_PARAMS

	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = errno.SUCCESS
		} else {
			code = errno.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code.Code,
		"msg":  code.Message,
		"data": data,
	})
}

// @Summary 获取多个文章
// @Produce  json
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500  {string} string "{"code":500,"msg":"fail"}"
// @Router /api/v1/articles [get]
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	var tagID int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagID = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagID

		valid.Min(tagID, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := errno.INVALID_PARAMS
	if !valid.HasErrors() {
		code = errno.SUCCESS

		data["lists"] = models.GetArticles(utils.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)

	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code.Code,
		"msg":  code.Message,
		"data": data,
	})
}

// @Summary 新增文章
// @Produce  json
// @Param tag_id body int true "TagID"
// @Param title body string true "Title"
// @Param desc body string true "Desc"
// @Param content body string true "Content"
// @Param created_by body string true "CreatedBy"
// @Param state body int true "State"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500  {string} string "{"code":500,"msg":"fail"}"
// @Router /api/v1/articles [post]
func AddArticle(c *gin.Context) {
	tagID := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagID, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := errno.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistTagById(tagID) {
			data := make(map[string]interface{})
			data["tag_id"] = tagID
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state

			models.AddArticle(data)
			code = errno.SUCCESS
		} else {
			code = errno.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code.Code,
		"msg":  code.Message,
		"data": make(map[string]interface{}),
	})
}

// @Summary 修改文章
// @Produce  json
// @Param id path int true "ID"
// @Param tag_id body string false "TagID"
// @Param title body string false "Title"
// @Param desc body string false "Desc"
// @Param content body string false "Content"
// @Param modified_by body string true "ModifiedBy"
// @Param state body int false "State"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500  {string} string "{"code":500,"msg":"fail"}"
// @Router /api/v1/articles/{id} [put]
func EditArticle(c *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	tagID := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := errno.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			if models.ExistTagById(tagID) {
				data := make(map[string]interface{})

				if tagID > 0 {
					data["tag_id"] = tagID
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}

				data["modified_by"] = modifiedBy

				models.EditArticle(id, data)
				code = errno.SUCCESS
			} else {
				code = errno.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = errno.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code.Code,
		"msg":  code.Message,
		"data": make(map[string]string),
	})
}

// @Summary 删除文章
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500  {string} string "{"code":500,"msg":"fail"}"
// @Router /api/v1/articles/{id} [delete]
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := errno.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
			code = errno.SUCCESS
		} else {
			code = errno.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code.Code,
		"msg":  code.Message,
		"data": make(map[string]string),
	})
}
