package dao

import (
	"blog_service/internal/model"
	"blog_service/pkg/app"
	"gorm.io/gorm"
)

type ITagDao interface {
	GetTag(tagID uint32, state uint8) (*model.Tag, error)
	Count(name string, state uint8) (int64, error)
	GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error)
	Create(name string, state uint8, createdBy string) error
	Update(id uint32, name string, state uint8, modifiedBy string) error
	Delete(id uint32) error
}
type TagDao struct {
	*Dao
}

func NewTag(engine *gorm.DB) ITagDao {
	return &TagDao{NewDao(engine)}
}
func (d *TagDao) GetTag(tagID uint32, state uint8) (*model.Tag, error) {
	var tag model.Tag
	err := d.engine.Where("id = ? AND state = ?", tagID, state).First(&tag).Error
	return &tag, err
}

func (d *TagDao) Count(name string, state uint8) (int64, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	return tag.Count(d.engine)
}

func (d *TagDao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *TagDao) Create(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	return tag.Create(d.engine)
}

func (d *TagDao) Update(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	value := map[string]interface{}{
		"name":        name,
		"state":       state,
		"modified_by": modifiedBy,
	}
	return tag.Update(d.engine, value)
}

func (d *TagDao) Delete(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}
