package blog

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BlogArticlesModel struct {
	Db   *gorm.DB
	Name string
}

type BlogArticle struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Cover       string    `json:"cover"`
	GoodNum     int       `json:"good_num"`
	View        int       `json:"view"`
	Recommended int       `json:"recommended"`
	Type        int       `json:"type"`
	Preface     string    `json:"preface"`
	HtmlTxtUrl  string    `json:"html_txt_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewBlogArticleModel(db *gorm.DB) *BlogArticlesModel {
	return &BlogArticlesModel{
		Db:   db,
		Name: "blog_articles",
	}
}

func (s *BlogArticlesModel) Create(in BlogArticle) (ret BlogArticle, err error) {
	err = s.Db.Table(s.Name).Create(&in).Error
	if err != nil {
		return
	}
	ret = in
	return
}

func (s *BlogArticlesModel) Update(in BlogArticle) (err error) {
	err = s.Db.Table(s.Name).Where("id = ?", in.Id).Updates(map[string]interface{}{
		"type":         in.Type,
		"preface":      in.Preface,
		"html_txt_url": in.HtmlTxtUrl,
		"name":         in.Name,
		"cover":        in.Cover,
	}).Error
	if err != nil {
		return
	}
	return
}

func (s *BlogArticlesModel) GetById(id int) (ret BlogArticle, err error) {
	err = s.Db.Table(s.Name).Where("id = ?", id).First(&ret).Error
	if err != nil {
		return
	}
	return
}

func (s *BlogArticlesModel) List(page, pageSize int) (ret []BlogArticle, total int, err error) {
	query := s.Db.Table(s.Name)
	err = query.Order("created_at desc").Offset(page * pageSize).Limit(pageSize).Find(&ret).Error
	if err != nil {
		return
	}
	err = query.Count(&total).Error
	return
}

type CreateBlogReq struct {
	Id      int    `json:"id"`
	Type    int    `json:"type"`
	Cover   string `json:"cover"`
	Desc    string `json:"desc"`
	HtmlTxt string `json:"html_txt"`
	MdTxt   string `json:"md_txt"`
	Name    string `json:"name"`
}

type ListRes struct {
	Total    int           `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
	List     []BlogArticle `json:"list"`
}
