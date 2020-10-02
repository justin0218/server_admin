package controllers
//
//import (
//	"github.com/astaxie/beego"
//	"io/ioutil"
//	"net/http"
//	"server_admin/initialize"
//	pb "server_admin/proto"
//	"github.com/micro/protobuf/proto"
//)
//
//type OpenApi struct {
//	beego.Controller
//}
//
//func (this *OpenApi) ReadFile() {
//	key := this.GetString("key")
//	resp := &pb.FileReadRes{}
//	res,err := initialize.RedisClient.Get(key).Bytes()
//	if err == nil{
//		resp.Txt = string(res)
//		respData,err := proto.Marshal(resp)
//		if err != nil{
//			this.Ctx.Output.Status = http.StatusBadRequest
//			return
//		}
//		this.Ctx.Output.Body(respData)
//		return
//	}
//	response,err := http.Get(key)
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusInternalServerError
//		return
//	}
//	defer response.Body.Close()
//	fdata,err := ioutil.ReadAll(response.Body)
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	if response.StatusCode == 200{
//		initialize.RedisClient.Set(key,fdata,-1)
//	}
//	resp.Txt = string(fdata)
//	respData,err := proto.Marshal(resp)
//	if err != nil{
//		this.Ctx.Output.Status = http.StatusBadRequest
//		return
//	}
//	this.Ctx.Output.Body(respData)
//	return
//}
//
//
