package services

import (
	"fmt"
	"server_admin/api"
	"server_admin/internal/models/bill"
	"server_admin/pkg/tool"
	"sort"
	"strconv"
	"strings"
)

type BillService struct {
}

func (s *BillService) Create(in bill.Bill) (ret bill.Bill, err error) {
	db := api.Mysql.Get()
	return bill.NewBillModel(db).Create(in)
}

func (s *BillService) SumBill(uid int) (ret []bill.SumBillData, err error) {
	db := api.Mysql.Get()
	list, _ := bill.NewBillModel(db).List(uid)
	sumMap := make(map[string]int)
	for _, v := range list {
		k := fmt.Sprintf("%d_%d", v.YearNum, v.MonthNum)
		sumMap[k] += v.Money
	}
	years := make([]int, 0)
	for k, v := range sumMap {
		item := bill.SumBillData{}
		dates := strings.Split(k, "_")
		item.Money = v
		item.YearNum, _ = strconv.Atoi(dates[0])
		item.MonthNum, _ = strconv.Atoi(dates[1])
		if tool.IntInArray(item.YearNum, years) == -1 {
			years = append(years, item.YearNum)
		}
		ret = append(ret, item)
	}
	sort.Ints(years)
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].YearNum < ret[j].YearNum
	})
	for dx, v := range ret {
		yearIdx := tool.IntInArray(v.YearNum, years)
		if yearIdx != -1 {
			ret[dx].Dx = yearIdx * 1000
		}
	}
	//for dx, v := range ret {
	//	ret[dx].Dx = v.Dx + v.MonthNum
	//}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Dx+ret[i].MonthNum < ret[j].Dx+ret[j].MonthNum
	})
	return
}
