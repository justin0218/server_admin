package controllers

import (
	"github.com/gin-gonic/gin"
	"server_admin/internal/models/bill"
	"server_admin/internal/services"
	"server_admin/pkg/jwt"
	"server_admin/pkg/resp"
	"time"
)

type BillController struct {
	billService *services.BillService
}

func (s *BillController) Create(c *gin.Context) {
	req := bill.CreateBillReq{}
	err := c.BindJSON(&req)
	if err != nil {
		resp.RespParamErr(c, err.Error())
		return
	}
	req.Uid = jwt.GetUid(c)
	if req.Money <= 0 {
		resp.RespParamErr(c)
		return
	}
	var t time.Time
	if req.Time == "" {
		t = time.Now().Local()
	}
	t, err = time.ParseInLocation("2006-01-02 15:04:05", req.Time, time.Local)
	if err != nil {
		resp.RespParamErr(c, err.Error())
		return
	}
	ret, err := s.billService.Create(bill.Bill{
		Note:     req.Note,
		Money:    req.Money,
		Time:     &t,
		YearNum:  t.Year(),
		MonthNum: int(t.Month()),
		DayNum:   t.Day(),
		Uid:      req.Uid,
	})
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}

func (s *BillController) SumBill(c *gin.Context) {
	ret, err := s.billService.SumBill(jwt.GetUid(c))
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}
