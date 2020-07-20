package model

type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
	State     uint8  `json:"state"`
}

func (ArticleTag) TableName() string {
	return "blog_article_tag"
}
