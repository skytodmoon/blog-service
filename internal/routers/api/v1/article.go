package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/internal/service"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/convert"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary 获取单个文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	param := service.ArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	// valid, errs := app.BindAndValid(c, &param)
	// if !valid {
	// 	response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	// 	return
	// }

	svc := service.New(c.Request.Context())
	article, err := svc.GetArticle(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}

	response.ToResponse(article)
	return
}
