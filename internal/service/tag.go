package service

import (
	"blog_service/global"
	dao2 "blog_service/internal/dao"
	"blog_service/internal/model"
	"blog_service/pkg/app"
	"context"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}
type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}
type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required.gte=1"`
}
type ITagService interface {
	CountTag(param *CountTagRequest) (int64, error)
	GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error)
	CreateTag(param *CreateTagRequest) error
	UpdateTag(param *UpdateTagRequest) error
	DeleteTag(param *DeleteTagRequest) error
}

type TagService struct {
	service Service
	dao     dao2.ITagDao
}

//tag 对应的service
func NewTagService(ctx context.Context) ITagService {
	return &TagService{service: New(ctx), dao: dao2.NewTag(global.DBEngine)}
}

func (s *TagService) CountTag(param *CountTagRequest) (int64, error) {
	return s.dao.Count(param.Name, param.State)
}

func (s *TagService) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return s.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (s *TagService) CreateTag(param *CreateTagRequest) error {
	return s.dao.Create(param.Name, param.State, param.CreatedBy)
}

func (s *TagService) UpdateTag(param *UpdateTagRequest) error {
	return s.dao.Update(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (s *TagService) DeleteTag(param *DeleteTagRequest) error {
	return s.dao.Delete(param.ID)
}
