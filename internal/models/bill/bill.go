package bill

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Bill struct {
	Uid      int        `json:"uid"`
	Note     string     `json:"note"`
	Money    int        `json:"money"`
	YearNum  int        `json:"year"`
	MonthNum int        `json:"month"`
	DayNum   int        `json:"day"`
	Time     *time.Time `json:"time"`
}

type BillModel struct {
	Db   *gorm.DB
	Name string
}

func NewBillModel(db *gorm.DB) *BillModel {
	return &BillModel{
		Db:   db,
		Name: "bills",
	}
}

func (s *BillModel) Create(in Bill) (ret Bill, err error) {
	err = s.Db.Table(s.Name).Create(&in).Error
	if err != nil {
		return
	}
	ret = in
	return
}

func (s *BillModel) List(uid int) (ret []Bill, err error) {
	err = s.Db.Table(s.Name).Where("uid = ?", uid).Find(&ret).Error
	if err != nil {
		return
	}
	return
}

type SumBillData struct {
	YearNum  int `json:"year"`
	MonthNum int `json:"month"`
	Money    int `json:"money"`
	Dx       int `json:"dx"`
}

type CreateBillReq struct {
	Uid   int    `json:"uid"`
	Time  string `json:"time"`
	Note  string `json:"note"`
	Money int    `json:"money"`
}
