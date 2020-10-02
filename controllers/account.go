package controllers

import (
	"github.com/astaxie/beego"
	"server_admin/models"
	"server_admin/initialize"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"server_admin/service"
	"time"
)

type Account struct {
	beego.Controller
	MakeBillReq models.MakeBillReq
	ChatData []models.ChatData
}

func (this *Account) AccountList() {
	initialize.O.Raw("select year_,month_,sum(money) as money from account_bills group by month_,year_ order by year_,month_").QueryRows(&this.ChatData)

	models.Success(this.Ctx,this.ChatData)
	return
}

func (this *Account) MakeBill() {
	request,err := ioutil.ReadAll(this.Ctx.Request.Body)
	if err != nil{
		models.Error(this.Ctx,err.Error(),http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(request,&this.MakeBillReq)
	if err != nil{
		models.Error(this.Ctx,err.Error(),http.StatusBadRequest)
		return
	}

	if this.MakeBillReq.Money <= 0 {
		models.Error(this.Ctx,"",http.StatusBadRequest)
		return
	}

	var t time.Time

	if this.MakeBillReq.Time == ""{
		t = time.Now().Local()
	}else{
		t,err = time.ParseInLocation("2006-01-02 15:04:05",this.MakeBillReq.Time,time.Local)
		if err != nil{
			models.Error(this.Ctx,err.Error(),http.StatusBadRequest)
			return
		}
	}

	_,err = initialize.O.Raw("insert account_bills set note = ?,money = ?,year_ = ?,month_ = ?,day_ = ?,time = ?",this.MakeBillReq.Note,this.MakeBillReq.Money,t.Year(),service.GetMonth(t),t.Day(),t).Exec()
	if err != nil{
		models.Error(this.Ctx,err.Error(),http.StatusBadRequest)
		return
	}

	models.Success(this.Ctx,1)
	return
}
