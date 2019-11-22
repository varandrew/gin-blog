package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/varandrew/gin-product/app/pkg/e"
)

// 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
}

// 获取多个文章
func GetArticles(c *gin.Context) {
}

// 新增文章
func AddArticle(c *gin.Context) {
}

// 修改文章
func EditArticle(c *gin.Context) {
}

// 删除文章
func DeleteArticle(c *gin.Context) {
}
