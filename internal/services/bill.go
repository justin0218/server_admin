package services

import (
	"server_admin/api"
	"server_admin/internal/models/bill"
)

type BillService struct {
}

func (s *BillService) Create(in bill.Bill) (ret bill.Bill, err error) {
	db := api.Mysql.Get()
	return bill.NewBillModel(db).Create(in)
}

func (s *BillService) SumBill() (ret []bill.SumBillData, err error) {
	db := api.Mysql.Get()
	return bill.NewBillModel(db).SumBill()
}
