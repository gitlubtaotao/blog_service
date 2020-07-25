package dao

import (
	"blog_service/internal/model"
	"gorm.io/gorm"
)

type IArticleTag interface {
	DeleteArticleTag(articleID uint32) error
	GetArticleTagByAID(articleID uint32) (model.ArticleTag, error)
	GetArticleTagListByTID(tagID uint32) ([]*model.ArticleTag, error)
	GetArticleTagListByAIDS(articleIDS []uint32) ([]*model.ArticleTag, error)
	CreateArticleTag(articleID, tagID uint32, createBy string) error
	UpdateArticleTag(articleID, tagID uint32, modifiedBy string) error
}
type ArticleTag struct {
	*Dao
}

func (a *ArticleTag) DeleteArticleTag(articleID uint32) error {
	return a.engine.Delete(&model.ArticleTag{}, "article_id = ? ", articleID).Error
}

func (a *ArticleTag) GetArticleTagByAID(articleID uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{ArticleID: articleID}
	return articleTag.GetByAID(a.engine)
}
func (a *ArticleTag) GetArticleTagListByTID(tagID uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{TagID: tagID}
	return articleTag.ListByTID(a.engine)
}
func (a *ArticleTag) GetArticleTagListByAIDS(articleIDS []uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{}
	return articleTag.ListByAIDs(a.engine, articleIDS)
}
func (a *ArticleTag) CreateArticleTag(articleID, tagID uint32, createBy string) error {
	articleTag := model.ArticleTag{
		Model:     &model.Model{CreatedBy: createBy},
		TagID:     tagID,
		ArticleID: articleID,
	}
	return articleTag.Create(a.engine)
}
func (a *ArticleTag) UpdateArticleTag(articleID, tagID uint32, modifiedBy string) error {
	articleTag := model.ArticleTag{ArticleID: articleID}
	values := map[string]interface{}{"article_id": articleID, "tag_id": tagID, "modified_by": modifiedBy}
	return articleTag.UpdateOne(a.engine, values)
}

func NewArticleTagDao(db *gorm.DB) IArticleTag {
	return &ArticleTag{NewDao(db)}
}
