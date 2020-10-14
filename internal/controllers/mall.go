package controllers

import (
	"github.com/gin-gonic/gin"
	"server_admin/internal/models/mall"
	"server_admin/internal/services"
	"server_admin/pkg/resp"
)

type MallController struct {
	goodsService services.GoodsService
}

func (s *MallController) CreateGoods(c *gin.Context) {
	req := mall.Goods{}
	err := c.BindJSON(&req)
	if err != nil {
		resp.RespParamErr(c)
		return
	}
	if req.Price <= 0 || req.Name == "" {
		resp.RespParamErr(c)
		return
	}
	_, err = s.goodsService.Create(req)
	if err != nil {
		resp.RespParamErr(c)
		return
	}
	resp.RespOk(c)
	return
}
