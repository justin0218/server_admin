package services

import (
	"github.com/google/uuid"
	"server_admin/api"
	"server_admin/internal/models/mall"
)

type GoodsService struct {
}

func (s *GoodsService) Create(in mall.Goods) (ret mall.Goods, err error) {
	db := api.Mysql.GetMall()
	in.Status = 1
	in.GoodsId = uuid.New().String()
	return mall.NewGoodsModel(db).Create(in)
}
