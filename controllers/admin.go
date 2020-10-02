package controllers

import (
	//"github.com/astaxie/beego"
	//"server_admin/util"
	//"io/ioutil"
	//"server_admin/models"
	//"encoding/json"
	//"net/http"
	//"server_admin/initialize"
	//"server_admin/service"
	//"github.com/google/uuid"
	//"fmt"
	//"strings"
	//"github.com/astaxie/beego/orm"
)

//type Admin struct {
//	beego.Controller
//	Encryption util.Encryption
//	LoginReq models.LoginReq
//	LoginRes models.LoginRes
//	TokenSvr service.Token
//	GetBlogListRes struct{
//		Total int `json:"total"`
//		Page int `json:"page"`
//		Size int `json:"size"`
//		List []models.GetBlogListRes `json:"list"`
//	}
//	DelBlogReq struct{
//		Id int `json:"id"`
//	}
//}
//
//func (this *Admin) Login(){
//	requestData,err := ioutil.ReadAll(this.Ctx.Request.Body)
//	if err != nil{
//		errors.JsonError(this.Ctx,errors.ErrParam)
//		return
//	}
//	err = json.Unmarshal(requestData,&this.LoginReq)
//	if err != nil{
//		errors.JsonError(this.Ctx,errors.ErrParam)
//		return
//	}
//	if this.LoginReq.Username == "" || this.LoginReq.Password == ""{
//		errors.JsonError(this.Ctx,errors.ErrParam)
//		return
//	}
//	password := ""
//	var uid int64
//	err = initialize.O.Raw("select uid,password from blog_admin_users where username = ?",this.LoginReq.Username).QueryRow(&uid,&password)
//	if err != nil{
//		errors.JsonError(this.Ctx,errors.ErrPasswd)
//		return
//	}
//	isTrue := this.Encryption.CheckMd5(this.LoginReq.Password,password)
//	if !isTrue{
//		errors.JsonError(this.Ctx,errors.ErrPasswd)
//		return
//	}
//	tokenStr,err := this.TokenSvr.CreateToken(uid)
//	if err != nil{
//		models.Error(this.Ctx,err.Error(),http.StatusBadRequest)
//		return
//	}
//	this.LoginRes.Uid = uid
//	this.LoginRes.Token = tokenStr
//	errors.JsonOK(this.Ctx,this.LoginRes)
//	return
//}
//
//func (this *Admin) CreateBlog() {
//	requestData,err := ioutil.ReadAll(this.Ctx.Request.Body)
//	if err != nil{
//		errors.JsonError(this.Ctx,errors.ErrParam.WithMessage(err.Error()))
//		return
//	}
//	createBlogReq := models.CreateBlogReq{}
//	err = json.Unmarshal(requestData,&createBlogReq)
//	if err != nil{
//		errors.JsonError(this.Ctx,errors.ErrParam.WithMessage(err.Error()))
//		return
//	}
//	if createBlogReq.Name == "" || createBlogReq.Type <= 0 || createBlogReq.Desc == "" || createBlogReq.HtmlTxt == "" || createBlogReq.MdTxt == ""{
//		errors.JsonError(this.Ctx,errors.ErrParam.WithMessage("参数错误"))
//		return
//	}
//
//	if createBlogReq.Id == 0{
//		fname := uuid.New().String()
//		mdName := fmt.Sprintf("md/%s.md",fname)
//		hTxtname := fmt.Sprintf("htxt/%s.shtml",fname)
//		service.UploadFile(mdName,[]byte(createBlogReq.MdTxt))
//		html_txt_url,err := service.UploadFile(hTxtname,[]byte(createBlogReq.HtmlTxt))
//		if err != nil{
//			errors.JsonError(this.Ctx,errors.ErrBase.WithMessage(err.Error()))
//			return
//		}
//		o := orm.NewOrm()
//		o.Begin()
//		lres,err := o.Raw("insert blog_articles set type = ?,preface = ?,html_txt_url = ?,name = ?,cover = ?",createBlogReq.Type,createBlogReq.Desc,html_txt_url,createBlogReq.Name,createBlogReq.Cover).Exec()
//		if err != nil{
//			o.Rollback()
//			errors.JsonError(this.Ctx,errors.ErrBase.WithMessage(err.Error()))
//			return
//		}
//		_,err = o.Raw("update blog_types set blog_num = blog_num + 1 where id = ?",createBlogReq.Type).Exec()
//		if err != nil{
//			o.Rollback()
//			errors.JsonError(this.Ctx,errors.ErrBase.WithMessage(err.Error()))
//			return
//		}
//		o.Commit()
//		lasterId,_ := lres.LastInsertId()
//		errors.JsonOK(this.Ctx,lasterId)
//		return
//	}
//
//	fname := uuid.New().String()
//	mdName := fmt.Sprintf("md/%s.md",fname)
//	hTxtname := fmt.Sprintf("htxt/%s.shtml",fname)
//	service.UploadFile(mdName,[]byte(createBlogReq.MdTxt))
//	html_txt_url,err := service.UploadFile(hTxtname,[]byte(createBlogReq.HtmlTxt))
//	if err != nil{
//		errors.JsonError(this.Ctx,errors.ErrBase.WithMessage(err.Error()))
//		return
//	}
//	tp := 0
//	initialize.O.Raw("select type from blog_articles where id = ?",createBlogReq.Id).QueryRow(&tp)
//	o := orm.NewOrm()
//	o.Begin()
//	if tp != createBlogReq.Type{
//		_,err = o.Raw("update blog_types set blog_num = blog_num - 1 where id = ?",tp).Exec()
//		if err != nil{
//			o.Rollback()
//			errors.JsonError(this.Ctx,errors.ErrBase.WithMessage(err.Error()))
//			return
//		}
//		_,err = initialize.O.Raw("update blog_types set blog_num = blog_num + 1 where id = ?",createBlogReq.Type).Exec()
//		if err != nil{
//			o.Rollback()
//			errors.JsonError(this.Ctx,errors.ErrBase.WithMessage(err.Error()))
//			return
//		}
//	}
//	_,err = o.Raw("update blog_articles set type = ?,preface = ?,html_txt_url = ?,name = ?,cover = ? where id = ?",createBlogReq.Type,createBlogReq.Desc,html_txt_url,createBlogReq.Name,createBlogReq.Cover,createBlogReq.Id).Exec()
//	if err != nil{
//		o.Rollback()
//		errors.JsonError(this.Ctx,errors.ErrBase.WithMessage(err.Error()))
//		return
//	}
//	o.Commit()
//	errors.JsonOK(this.Ctx,createBlogReq.Id)
//	return
//}
//
//func (this *Admin) UploadFile() {
//	//fmt.Println(this.Ctx.)
//	f,h,err := this.GetFile("editormd-image-file")
//	if err != nil{
//		models.Error(this.Ctx,err.Error(),http.StatusBadRequest)
//		return
//	}
//	defer f.Close()
//
//	path := this.GetString("path","images")
//
//	flen := strings.Split(h.Filename,".")
//	if len(flen) != 2 {
//		models.Error(this.Ctx,"格式不正确",http.StatusBadRequest)
//		return
//	}
//	suffix := flen[1]
//	fname := uuid.New().String()
//	obj := fmt.Sprintf("%s/%s.%s",path,fname,suffix)
//	fileBytes,err := ioutil.ReadAll(f)
//	if err != nil{
//		models.Error(this.Ctx,err.Error(),http.StatusBadRequest)
//		return
//	}
//	url,err := service.UploadFile(obj,fileBytes)
//	if err != nil{
//		models.Error(this.Ctx,err.Error(),http.StatusBadRequest)
//		return
//	}
//	mapRes := make(map[string]string)
//	mapRes["url"] = url
//	mapRes["success"] = "1"
//	this.Ctx.Output.JSON(mapRes,false,false)
//	return
//}
//
//func (this *Admin) GetBlogList() {
//	page,err := this.GetInt("page",1)
//	if err != nil{
//		models.Error(this.Ctx,err.Error(),http.StatusBadRequest)
//		return
//	}
//	tp,err := this.GetInt("type",-1)
//	if err != nil{
//		models.Error(this.Ctx,err.Error(),http.StatusBadRequest)
//		return
//	}
//
//	num := 50
//	page--
//	tpEqual := "="
//	if tp == -1{
//		tpEqual = "!="
//	}
//
//	initialize.O.Raw("select * from blog_articles where type "+tpEqual+" ? order by update_time desc limit ?,?",tp,page*num,num).QueryRows(&this.GetBlogListRes.List)
//	initialize.O.Raw("select count(id) from blog_articles where type "+tpEqual+" ?",tp).QueryRow(&this.GetBlogListRes.Total)
//	this.GetBlogListRes.Size = num
//	this.GetBlogListRes.Page = page + 1
//	errors.JsonOK(this.Ctx,this.GetBlogListRes)
//	return
//}
//
//func (this *Admin) GetBlogDetail() {
//	id,err := this.GetInt("id")
//	if err != nil{
//		errors.JsonError(this.Ctx,errors.ErrParam)
//		return
//	}
//	res := models.GetBlogListRes{}
//	initialize.O.Raw("select * from blog_articles where id = ?",id).QueryRow(&res)
//	errors.JsonOK(this.Ctx,res)
//	return
//}
//
//func (this *Admin) DelBlog() {
//
//	//bd,err := ioutil.ReadAll(this.Ctx.Request.Body)
//	//if err != nil{
//	//	models.Error(this.Ctx,err.Error(),http.StatusBadRequest)
//	//	return
//	//}
//	//err = json.Unmarshal(bd,&this.DelBlogReq)
//	//if err != nil{
//	//	models.Error(this.Ctx,err.Error(),http.StatusBadRequest)
//	//	return
//	//}
//	//initialize.O.Raw("delete from blog_articles where id = ?",this.DelBlogReq.Id).Exec()
//	//models.Success(this.Ctx,this.GetBlogDetailRes)
//	//return
//}
//
//func (this *Admin) ReadFile() {
//	key := this.GetString("key")
//	res,err := initialize.RedisClient.Get(key).Result()
//	if err == nil{
//		errors.JsonOK(this.Ctx,res)
//		return
//	}
//	response,err := http.Get(key)
//	if err != nil{
//		errors.JsonError(this.Ctx,errors.ErrBase.WithMessage(err.Error()))
//		return
//	}
//	defer response.Body.Close()
//	fdata,err := ioutil.ReadAll(response.Body)
//	if err != nil{
//		errors.JsonError(this.Ctx,errors.ErrBase.WithMessage(err.Error()))
//		return
//	}
//	initialize.RedisClient.Set(key,fdata,-1)
//	errors.JsonOK(this.Ctx,string(fdata))
//	return
//	return
//}





