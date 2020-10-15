package mall

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Goods struct {
	GoodsId   string    `json:"goods_id"`
	Uid       int       `json:"uid"`
	Name      string    `json:"name"`
	Details   string    `json:"details"`
	Price     int       `json:"price"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GoodsModel struct {
	Db   *gorm.DB
	Name string
}

func NewGoodsModel(db *gorm.DB) *GoodsModel {
	return &GoodsModel{
		Db:   db,
		Name: "goods",
	}
}

func (s *GoodsModel) Create(in Goods) (ret Goods, err error) {
	err = s.Db.Table(s.Name).Create(&in).Error
	if err != nil {
		return
	}
	ret = in
	return
}
