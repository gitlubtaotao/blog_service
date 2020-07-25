package dao

import (
	"blog_service/internal/model"
	"blog_service/pkg/app"
	"gorm.io/gorm"
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

type IArticle interface {
	CountArticleListByTagID(id uint32, state uint8) (int64, error)
	GetArticleListByTagID(id uint32, state uint8, page, pageSize int) ([]*model.ArticleRow, error)
	DeleteArticle(id uint32) error
	GetArticle(id uint32, state uint8) (model.Article, error)
	CreateArticle(param *Article) (*model.Article, error)
	UpdateArticle(id uint32, value map[string]interface{}) error
}

type ArticleDao struct {
	*Dao
}

func (a *ArticleDao) CountArticleListByTagID(id uint32, state uint8) (int64, error) {
	article := model.Article{State: state}
	return article.CountByTagID(a.engine, id)
}

func (a *ArticleDao) GetArticleListByTagID(id uint32, state uint8, page, pageSize int) ([]*model.ArticleRow, error) {
	article := model.Article{State: state}
	return article.ListByTagID(a.engine, id, app.GetPageOffset(page, pageSize), pageSize)
}

func (a *ArticleDao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(a.engine)
}

func (a *ArticleDao) GetArticle(id uint32, state uint8) (model.Article, error) {
	article := model.Article{Model: &model.Model{ID: id}, State: state}
	return article.Get(a.engine)
}

func (a *ArticleDao) UpdateArticle(id uint32, values map[string]interface{}) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Update(a.engine, values)
}

func (a *ArticleDao) CreateArticle(param *Article) (*model.Article, error) {
	article := model.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Model:         &model.Model{CreatedBy: param.CreatedBy},
	}
	return article.Create(a.engine)
}

func NewArticleDao(engine *gorm.DB) IArticle {
	return &ArticleDao{NewDao(engine)}
}
