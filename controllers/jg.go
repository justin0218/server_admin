package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"server_admin/models"
	"server_admin/service"
	uuid2 "github.com/google/uuid"
	"server_admin/initialize"
	"fmt"
)

type JgController struct {
	beego.Controller
}

func (this *JgController) CreateAppointment() {
	reqDataByte,err := ioutil.ReadAll(this.Ctx.Request.Body)
	if err != nil{
		this.Ctx.Output.Status = http.StatusBadRequest
		return
	}
	req := models.JgAppointment{}
	err = json.Unmarshal(reqDataByte,&req)
	if err != nil{
		fmt.Println(err.Error())
		this.Ctx.Output.Status = http.StatusBadRequest
		return
	}
	_,err = service.BaseSvr.InserSql("jg_appointments",req)
	if err != nil{
		fmt.Println(err.Error())
		this.Ctx.Output.Status = http.StatusBadRequest
		return
	}
	this.Ctx.Output.Status = http.StatusOK
}

func (this *JgController) GetAppointmentList(){
	uuid := uuid2.New().String()
	name := this.GetString("name",uuid)
	sex,err := this.GetInt("sex",0)
	if err != nil{
		this.Ctx.Output.Status = http.StatusBadRequest
		return
	}
	page,err := this.GetInt("page",1)
	if err != nil{
		this.Ctx.Output.Status = http.StatusBadRequest
		return
	}
	size,err := this.GetInt("size",20)
	if err != nil{
		this.Ctx.Output.Status = http.StatusBadRequest
		return
	}
	timeBucket,err := this.GetInt("time_bucket",0)
	if err != nil{
		this.Ctx.Output.Status = http.StatusBadRequest
		return
	}
	checkMethod,err := this.GetInt("check_method",0)
	if err != nil{
		this.Ctx.Output.Status = http.StatusBadRequest
		return
	}
	checkProject,err := this.GetInt("check_project",0)
	if err != nil{
		this.Ctx.Output.Status = http.StatusBadRequest
		return
	}

	phone := this.GetString("phone",uuid)
	appointmentDate := this.GetString("appointment_date",uuid)
	nameEq,sexEq,timeBucketEq,checkMethodEq,phoneEq,appointmentDateEq,checkProjectEq := "like","=","=","=","like","=","="
	if name == uuid{
		nameEq = "not like"
	}
	if sex == 0{
		sexEq = "!="
	}
	if timeBucket == 0{
		timeBucketEq = "!="
	}
	if checkMethod == 0{
		checkMethodEq = "!="
	}
	if checkProject == 0{
		checkProjectEq = "!="
	}
	if phone == uuid{
		phoneEq = "not like"
	}
	if appointmentDate == uuid{
		appointmentDateEq = "!="
	}
	page--
	res := models.GetAppointmentListRes{}
	listSql := fmt.Sprintf(`
		select * from jg_appointments
		where name %s concat('%%',?,'%%')
		and sex %s ?
		and time_bucket %s ?
		and check_method %s ?
		and phone %s concat('%%',?,'%%')
		and appointment_date %s ?
		and check_project %s ?
		order by appointment_date
		limit ?,?
	`,nameEq,sexEq,timeBucketEq,checkMethodEq,phoneEq,appointmentDateEq,checkProjectEq)
	countSql := fmt.Sprintf(`
		select count(*) from jg_appointments
		where name %s concat('%%',?,'%%')
		and sex %s ?
		and time_bucket %s ?
		and check_method %s ?
		and phone %s concat('%%',?,'%%')
		and appointment_date %s ?
		and check_project %s ?
	`,nameEq,sexEq,timeBucketEq,checkMethodEq,phoneEq,appointmentDateEq,checkProjectEq)
	initialize.O.Raw(listSql,name,sex,timeBucket,checkMethod,phone,appointmentDate,checkProject,page*size,size).QueryRows(&res.List)
	initialize.O.Raw(countSql,name,sex,timeBucket,checkMethod,phone,appointmentDate,checkProject).QueryRow(&res.Total)
	res.Size = size
	res.Page++
	this.Ctx.Output.JSON(res,false,false)
}


