package model

import (
	"blog_service/pkg/app"
	"gorm.io/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (Tag) TableName() string {
	return "blog_tags"
}

// 统计tag的数量
func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var (
		tags []*Tag
		err  error
	)
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("status = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, err
}

func (t Tag) Create(db *gorm.DB) error  {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB,value map[string]interface{}) error  {
	return db.Model(&t).Where("id = ? and is_del = ?",t.ID,0).Updates(value).Error
}

func (t Tag) Delete(db *gorm.DB) error  {
	return db.Where("id = ? and is_del = ?",t.ID,0).Delete(&t).Error
}
