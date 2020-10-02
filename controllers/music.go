package controllers

import (
	"github.com/astaxie/beego"
	"server_admin/initialize"
	"server_admin/models"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/google/uuid"
	"server_admin/service"
)

type MusicController struct {
	beego.Controller
}




func (this *MusicController) GetMusicList() {
	res := []models.Music{}
	initialize.O.Raw("select * from songs order by create_time").QueryRows(&res)
	models.NSuccess(this.Ctx,res)
	return
}

func (this *MusicController) CreateMusic() {
	reqDataByte,err := ioutil.ReadAll(this.Ctx.Request.Body)
	if err != nil{
		models.NError(this.Ctx,err.Error())
		return
	}
	req := models.CreateMusicReq{}
	json.Unmarshal(reqDataByte,&req)
	if req.Name == "" || req.Singer == "" || req.Url == ""{
		models.NError(this.Ctx,"参数错误")
		return
	}
	if req.Type == 1{ //直接入库
		initialize.O.Raw("insert songs set name = ?,url = ?,singer = ?",req.Name,req.Url,req.Singer).Exec()
		models.NSuccess(this.Ctx,1)
		return
	}

	resp,err := http.Get(req.Url)
	if err != nil{
		models.NError(this.Ctx,err.Error())
		return
	}
	defer resp.Body.Close()
	fname := uuid.New().String()
	obj := fmt.Sprintf("songs/%s.mp3",fname)
	fileByte,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		models.Error(this.Ctx,err.Error(),http.StatusBadRequest)
		return
	}
	furl,err := service.UploadFile(obj,fileByte)
	if err != nil{
		models.Error(this.Ctx,err.Error(),http.StatusBadRequest)
		return
	}
	initialize.O.Raw("insert songs set name = ?,url = ?,singer = ?",req.Name,furl,req.Singer).Exec()
	models.NSuccess(this.Ctx,1)
	return
}


