package dao

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
)

type Article struct {
	ID            uint32 `json:"id"`
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

func (d *Dao) GetArticle(id uint32, state uint8) (model.Article, error) {
	article := model.Article{Model: &model.Model{ID: id}, State: state}
	return article.Get(d.engine)
}
