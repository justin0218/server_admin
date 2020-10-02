package blog

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BlogTypeModel struct {
	Db   *gorm.DB
	Name string
}

type BlogType struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	GoodNum     int       `json:"good_num"`
	View        int       `json:"view"`
	Recommended int       `json:"recommended"`
	Type        int       `json:"type"`
	Preface     int       `json:"preface"`
	HtmlTxtUrl  string    `json:"html_txt_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewBlogTypeModel(db *gorm.DB) *BlogTypeModel {
	return &BlogTypeModel{
		Db:   db,
		Name: "blog_types",
	}
}

func (s *BlogTypeModel) IncrBlogNum(id int) (err error) {
	err = s.Db.Table(s.Name).Where("id = ?", id).UpdateColumn("blog_num", gorm.Expr("blog_num + ?", 1)).Error
	if err != nil {
		return
	}
	return
}

func (s *BlogTypeModel) DecrBlogNum(id int) (err error) {
	err = s.Db.Table(s.Name).Where("id = ?", id).UpdateColumn("blog_num", gorm.Expr("blog_num - ?", 1)).Error
	if err != nil {
		return
	}
	return
}
