package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	article := v1.NewArticle()
	apiv1 := r.Group("/api/v1")
	apiv1.Use() //middleware.JWT()
	{
		// 获取指定文章
		apiv1.GET("/articles/:id", article.Get)

	}

	return r
}
