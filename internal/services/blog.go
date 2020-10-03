package services

import (
	"server_admin/api"
	"server_admin/internal/models/blog"
)

type BlogService struct {
}

func (s *BlogService) Create(in blog.BlogArticle) (ret blog.BlogArticle, err error) {
	db := api.Mysql.Get()
	tx := db.Begin()
	ret, err = blog.NewBlogArticleModel(tx).Create(in)
	if err != nil {
		tx.Rollback()
		return
	}
	err = blog.NewBlogTypeModel(tx).IncrBlogNum(in.Type)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}

func (s *BlogService) Update(in blog.BlogArticle) (err error) {
	db := api.Mysql.Get()
	tx := db.Begin()
	blogDetail, e := blog.NewBlogArticleModel(tx).GetById(in.Id)
	if e != nil {
		err = e
		return
	}
	if in.Type != blogDetail.Type {
		err = blog.NewBlogTypeModel(tx).DecrBlogNum(blogDetail.Type)
		if err != nil {
			tx.Rollback()
			return
		}
		err = blog.NewBlogTypeModel(tx).IncrBlogNum(in.Type)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	err = blog.NewBlogArticleModel(tx).Update(in)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	return
}

func (s *BlogService) List(page, pageSize, uid int) (ret blog.ListRes, err error) {
	db := api.Mysql.Get()
	page--
	ret.List, ret.Total, err = blog.NewBlogArticleModel(db).List(page, pageSize, uid)
	ret.Page = page + 1
	ret.PageSize = pageSize
	return
}

func (s *BlogService) Detail(id int) (ret blog.BlogArticle, err error) {
	db := api.Mysql.Get()
	ret, err = blog.NewBlogArticleModel(db).GetById(id)
	return
}
