package controllers
//
//import (
//	"github.com/astaxie/beego"
//	"server_admin/models"
//	"server_admin/initialize"
//	"net/http"
//	"server_admin/service"
//	"fmt"
//	"time"
//	"io/ioutil"
//	pb "server_admin/proto"
//	"github.com/micro/protobuf/proto"
//	"github.com/google/uuid"
//)
//
//type Blog struct {
//	beego.Controller
//}
//
//func (this *Blog) Getrecommend() {
//	res := &pb.BlogListRes{}
//	initialize.O.Raw("select * from blog_articles where recommended = 1 and id != 35").QueryRows(&res.List)
//	out,err := proto.Marshal(res)
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	this.Ctx.Output.Body(out)
//	return
//}
//
//func (this *Blog) GetBlogRanking() {
//	limit,err := this.GetInt("limit")
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	res := &pb.BlogListRes{}
//	initialize.O.Raw("select * from blog_articles where id != 35 order by view desc limit ?",limit).QueryRows(&res.List)
//	out,err := proto.Marshal(res)
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	this.Ctx.Output.Body(out)
//	return
//}
//
//func (this *Blog) GetBlogTypes() {
//	tps := pb.Tps{}
//	initialize.O.Raw("select * from blog_types").QueryRows(&tps.List)
//	out,err := proto.Marshal(&tps)
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	this.Ctx.Output.Body(out)
//	return
//}
//
//func (this *Blog) GetBlogList() {
//	tp,err := this.GetInt("tp",-1)
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	tpEq,nameEq := "=","like"
//	if tp == -1{
//		tpEq = "!="
//	}
//	uuname := uuid.New().String()
//	name := this.GetString("name",uuname)
//	if name == uuname{
//		nameEq = "not like"
//	}
//	res := &pb.BlogListRes{}
//	_,err = initialize.O.Raw("select * from blog_articles where name "+nameEq+" concat('%',?,'%') and type "+tpEq+" ? and id != 35 order by recommended desc",name,tp).QueryRows(&res.List)
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	out,err := proto.Marshal(res)
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	this.Ctx.Output.Body(out)
//	return
//}
//
//
//func (this *Blog) GetBlogDetail() {
//	id,err := this.GetInt("id")
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	res := &pb.DetailRes{}
//	initialize.O.Raw("select * from blog_articles where id = ?",id).QueryRow(&res)
//	initialize.O.Raw("select * from blog_articles where id > ? and id != 35 limit 1",id).QueryRow(&res.Next)
//	initialize.O.Raw("select * from blog_articles where id < ? and id != 35 order by id desc limit 1",id).QueryRow(&res.Prev)
//	resdata,err := proto.Marshal(res)
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	go func() {
//		ip := this.Ctx.Input.IP()
//		ipdetail := service.GetIpDetailFromTaobao(ip)
//		_,err := initialize.O.Raw("insert blog_views set ip = ?,country = ?,region = ?,city = ?,isp = ?,blog_id = ?",ip,ipdetail.Data.Country,ipdetail.Data.Region,ipdetail.Data.City,ipdetail.Data.Isp,id).Exec()
//		if err == nil{
//			initialize.O.Raw("update blog_articles set view = view + 1 where id = ?",id).Exec()
//		}
//	}()
//	this.Ctx.Output.Body(resdata)
//	return
//}
//
//func (this *Blog) SubmitMessage(){
//	data,err := ioutil.ReadAll(this.Ctx.Request.Body)
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	req := &pb.BlogComment{}
//	err = proto.Unmarshal(data,req)
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	if req.Content == ""{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	ip := this.Ctx.Input.IP()
//	key := fmt.Sprintf("%s:%s:%d",time.Now().Format("2006-01-02"),ip,req.BlogId)
//	num,err := initialize.RedisClient.Incr(key).Result()
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusInternalServerError
//		return
//	}
//	if num > 10{
//		this.Ctx.Output.Status = http.StatusForbidden
//		return
//	}
//	tempUserName := fmt.Sprintf("IP-%s-用户",ip)
//	_,err = initialize.O.Raw("insert blog_comments set name = ?,content = ?,blog_id = ?",tempUserName,req.Content,req.BlogId).Exec()
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusInternalServerError
//		return
//	}
//	this.Ctx.Output.Status = http.StatusOK
//	return
//}
//
//func (this *Blog) GetVerifyCode() {
//	verifyKey,val := service.CodeCaptchaCreate()
//	mapRes := make(map[string]string)
//	mapRes["id_key"] = verifyKey
//	mapRes["value"] = val
//	models.Success(this.Ctx,mapRes)
//	return
//}
//
//func (this *Blog) GetMessageList() {
//	blogId,err := this.GetInt("blog_id",0)
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	res := &pb.BlogComments{}
//	initialize.O.Raw("select * from blog_comments where blog_id = ? limit 100",blogId).QueryRows(&res.List)
//	initialize.O.Raw("select count(id) from blog_comments where blog_id = ?",blogId).QueryRow(&res.Total)
//	out,err := proto.Marshal(res)
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	this.Ctx.Output.Body(out)
//	return
//}
//
//func (this *Blog) MakeGood() {
//	blogId,err := this.GetInt("blog_id",0)
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	go initialize.O.Raw("update blog_articles set good_num = good_num + 1 where id = ?",blogId).Exec()
//	this.Ctx.Output.Status = http.StatusOK
//	return
//}
//
//
//
