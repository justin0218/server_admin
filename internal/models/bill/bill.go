package bill

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Bill struct {
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

func (s *BillModel) List() (ret []Bill, err error) {
	err = s.Db.Table(s.Name).Find(&ret).Error
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
	Time  string `json:"time"`
	Note  string `json:"note"`
	Money int    `json:"money"`
}
