package bill

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Bill struct {
	Note   string     `json:"note"`
	Money  int        `json:"money"`
	Year_  int        `json:"year_"`
	Month_ int        `json:"month_"`
	Day_   int        `json:"day_"`
	Time   *time.Time `json:"time"`
}

type BillModel struct {
	Db   *gorm.DB
	Name string
}

func NewBillModel(db *gorm.DB) *BillModel {
	return &BillModel{
		Db:   db,
		Name: "account_bills",
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

func (s *BillModel) SumBill() (ret []SumBillData, err error) {
	err = s.Db.Table(s.Name).Exec("select year_,month_,sum(money) as money from account_bills").Group("year_").Group("month_").Order("year_").Order("month_").Find(&ret).Error
	if err != nil {
		return
	}
	return
}

type SumBillData struct {
	Year_  int   `json:"year"`
	Month_ int   `json:"month"`
	Money  int64 `json:"money"`
}

type CreateBillReq struct {
	Time  string `json:"time"`
	Note  string `json:"note"`
	Money int    `json:"money"`
}
