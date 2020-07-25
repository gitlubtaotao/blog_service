package model

import (
	"blog_service/pkg/app"
	"gorm.io/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

type ArticleRow struct {
	ArticleID       uint32 `json:"article_id"`
	TagID           uint32 `json:"tag_id"`
	TagName         string `json:"tag_name"`
	ArticleTitle    string `json:"article_title"`
	ArticleDesc     string `json:"article_desc"`
	ArticleState    uint8  `json:"article_state"`
	ArticleImageUrl string `json:"article_image_url"`
	Content         string `json:"content"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (Article) TableName() string {
	return "blog_articles"
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (a Article) Update(db *gorm.DB, value interface{}) error {
	if err := db.Model(&a).Updates(value).Where("id = ? and is_del = ?", a.ID, 0).Error; err != nil {
		return err
	}
	return nil
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	db = db.Where("id = ? and state = ? and is_del = ?", a.Model.ID, a.State, 0)
	err := db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}
	return article, nil
}

func (a Article) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? and is_del = ?", a.Model.ID, 0).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{
		"blog_articles.id AS article_id",
		"blog_articles.title AS article_title",
		"blog_articles.desc AS article_desc",
		"blog_articles.cover_image_url AS article_image_url",
		"blog_articles.content AS content",
		"blog_articles.state as article_state",
		"tags.id AS tag_id",
		"tags.name AS tag_name",
	}
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	rows, err := db.Select(fields).Scopes(a.defaultScope).
		Where("at.tag_id = ? and ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var articles []*ArticleRow
	for rows.Next() {
		r := &ArticleRow{}
		if err := rows.Scan(&r); err != nil {
			return nil, err
		}
		articles = append(articles, r)
	}
	return articles, nil
}

func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (int64, error) {
	var count int64
	err := db.Scopes(a.defaultScope).
		Where("at.tag_id = ? and ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}

func (a Article) defaultScope(db *gorm.DB) *gorm.DB {
	return db.Table(ArticleTag{}.TableName() + " AS at").
		Joins("LEFT JOIN " + Tag{}.TableName() + " AS t ON at.tag_id = t.id  ").
		Joins("LEFT JOIN " + Article{}.TableName() + " AS ar ON at._article_id = ar.id")
	
}
