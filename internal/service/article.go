package service

import (
	"blog_service/global"
	dao2 "blog_service/internal/dao"
	"blog_service/internal/model"
	"blog_service/pkg/app"
	"context"
)

type Article struct {
	ID            uint32     `json:"id"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	Content       string     `json:"content"`
	CoverImageUrl string     `json:"cover_image_url"`
	State         uint8      `json:"state"`
	Tag           *model.Tag `json:"tag"`
}

type CountArticleRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}
type ArticleListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
	TagID uint32 `form:"id" binding:"required,gte=1"`
}
type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=3,max=100"`
	Desc          string `form:"desc" binding:"required,min=3,max=255"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	Content       string `form:"content"`
	CoverImageUrl string `form:"cover_image_url"`
	TagID         uint32 `form:"tag_id" binding:"gte=1"`
	CreatedBy     string `form:"created_by" binding:"required,min=3,max=100"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"min=3,max=100"`
	Desc          string `form:"desc" binding:"min=3,max= 255"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	Content       string `form:"content"`
	CoverImageUrl string `form:"cover_image_url"`
	TagID         uint32 `form:"tag_id" binding:"gte=1"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=3,max=100"`
}

type ArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type IArticleService interface {
	GetArticle(param *ArticleRequest) (*Article, error)
	GetArticleList(param *ArticleListRequest, page *app.Pager) ([]*Article, int64, error)
	CreateArticle(param *CreateArticleRequest) error
	UpdateArticle(param *UpdateArticleRequest) error
	DeleteArticle(param *DeleteRequest) error
}
type ArticleService struct {
	service       Service
	dao           dao2.IArticle
	articleTagDao dao2.IArticleTag
	tagDao        dao2.ITagDao
}

func (a *ArticleService) GetArticle(param *ArticleRequest) (*Article, error) {
	article, err := a.dao.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}
	articleTag, err := a.articleTagDao.GetArticleTagByAID(param.ID)
	if err != nil {
		return nil, err
	}
	tag, err := a.tagDao.GetTag(articleTag.ID, model.StateOpen)
	if err != nil {
		return nil, err
	}
	return &Article{
		ID:            article.ID,
		Title:         article.Title,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State:         article.State,
		Tag:           tag,
	}, nil
}

func (a *ArticleService) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*Article, int64, error) {
	articleCount, err := a.dao.CountArticleListByTagID(param.TagID, param.State)
	if err != nil {
		return nil, 0, err
	}
	articles, err := a.dao.GetArticleListByTagID(param.TagID, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}
	var articleList []*Article
	for _, article := range articles {
		articleList = append(articleList, &Article{
			Title:         article.ArticleTitle,
			Desc:          article.ArticleDesc,
			Content:       article.Content,
			CoverImageUrl: article.ArticleImageUrl,
			State:         article.ArticleState,
			Tag: &model.Tag{
				Name:  article.TagName,
				Model: &model.Model{ID: article.ArticleID},
			},
		})
	}
	return articleList, articleCount, nil
}

func (a *ArticleService) CreateArticle(param *CreateArticleRequest) error {
	//TODO CreateArticle and CreateArticleTag 应该在一个事务里面进行，这个我认为是一个bug
	article, err := a.dao.CreateArticle(&dao2.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		CreatedBy:     param.CreatedBy,
		State:         param.State,
	})
	if err != nil {
		return err
	}
	if param.TagID != 0 {
		err = a.articleTagDao.CreateArticleTag(article.ID, param.TagID, param.CreatedBy)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *ArticleService) UpdateArticle(param *UpdateArticleRequest) error {
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}
	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.CoverImageUrl
	}
	if param.Desc != "" {
		values["desc"] = param.Desc
	}
	if param.Content != "" {
		values["content"] = param.Content
	}
	//TODO updateArticle and UpdateArticleTag 应该在一个事务里面进行，这个我认为是一个bug
	if err := a.dao.UpdateArticle(param.ID, values); err != nil {
		return err
	}
	err := a.articleTagDao.UpdateArticleTag(param.ID, param.TagID, param.ModifiedBy)
	return err
}

func (a *ArticleService) DeleteArticle(param *DeleteRequest) error {
	//TODO DeleteArticle and DeleteArticleTag 应该在一个事务里面进行，这个我认为是一个bug
	err := a.dao.DeleteArticle(param.ID)
	if err != nil {
		return nil
	}
	
	err = a.articleTagDao.DeleteArticleTag(param.ID)
	return err
}

//tag 对应的service
func NewArticleService(ctx context.Context) IArticleService {
	return &ArticleService{service: New(ctx),
		dao:           dao2.NewArticleDao(global.DBEngine),
		articleTagDao: dao2.NewArticleTagDao(global.DBEngine),
		tagDao:        dao2.NewTag(global.DBEngine),
	}
}
