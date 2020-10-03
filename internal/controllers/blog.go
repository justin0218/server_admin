package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server_admin/internal/models/blog"
	"server_admin/internal/services"
	"server_admin/pkg/jwt"
	"server_admin/pkg/resp"
	"server_admin/pkg/upload"
	"strconv"
)

type BlogController struct {
	blogService *services.BlogService
}

func (s *BlogController) CreateBlog(c *gin.Context) {
	req := blog.CreateBlogReq{}
	err := c.BindJSON(&req)
	if err != nil {
		resp.RespParamErr(c)
		return
	}
	if req.Name == "" || req.Type <= 0 || req.Desc == "" || req.HtmlTxt == "" || req.MdTxt == "" {
		resp.RespParamErr(c)
		return
	}
	fname := uuid.New().String()
	mdName := fmt.Sprintf("md/%s.md", fname)
	hTxtname := fmt.Sprintf("htxt/%s.shtml", fname)
	_, err = upload.UploadFile(mdName, []byte(req.MdTxt))
	if err != nil {
		resp.RespInternalErr(c)
		return
	}
	htmlTxtUrl, err := upload.UploadFile(hTxtname, []byte(req.HtmlTxt))
	if err != nil {
		resp.RespInternalErr(c)
		return
	}
	in := blog.BlogArticle{
		Type:       req.Type,
		Preface:    req.Desc,
		HtmlTxtUrl: htmlTxtUrl,
		Name:       req.Name,
		Cover:      req.Cover,
		Uid:        jwt.GetUid(c),
	}
	if req.Id == 0 {
		ret, err := s.blogService.Create(in)
		if err != nil {
			resp.RespInternalErr(c)
			return
		}
		resp.RespOk(c, ret)
		return
	}
	in.Id = req.Id
	err = s.blogService.Update(in)
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, in)
	return
}

func (s *BlogController) GetBlogList(c *gin.Context) {
	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	pageSize := c.Query("page_size")
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		pageSizeInt = 20
	}
	if pageSizeInt > 500 {
		resp.RespParamErr(c, "分页最大不能超过500")
		return
	}
	ret, err := s.blogService.List(pageInt, pageSizeInt, jwt.GetUid(c))
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}

func (s *BlogController) Detail(c *gin.Context) {
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		resp.RespParamErr(c)
		return
	}
	ret, err := s.blogService.Detail(idInt)
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}
