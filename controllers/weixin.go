package controllers
//
//import (
//	"github.com/astaxie/beego"
//	"server_admin/service"
//	"server_admin/models"
//	"net/http"
//	"fmt"
//	"io/ioutil"
//	"encoding/json"
//	"server_admin/initialize"
//	"github.com/astaxie/beego/orm"
//	"time"
//	"strconv"
//)
//
//type Weixin struct {
//	beego.Controller
//	UserInfo user.Info
//	AuthRes models.AuthRes `json:"auth_res"`
//}
//
//type Tests struct {
//	A string `json:"a"`
//	B string `json:"b"`
//	C bool `json:"c"`
//}
//
//
//func (this *Weixin) Test() {
//	this.Ctx.Redirect(301,`http://www.momoman.cn`)
//	return
//}
//
//func (this *Weixin) TestFq() {
//
//	openid := this.GetString("openid","oBYAkw3URP9pAQekMZ1GYmuNfFfQ")
//	am := this.GetString("am","100")
//
//	appid := "wx19f2c0be339ede12"
//	mcId := "1493165212"
//
//	par := strconv.FormatInt(time.Now().UnixNano(),10)
//	desc := "wcaooo"
//	cliIp := "140.143.188.219"
//	a,b,c := service.WeixinSvr.WithdrawMoney(appid,mcId,openid,am,par,desc,cliIp)
//	dd := Tests{a,b,c}
//	models.Success(this.Ctx,dd)
//}
//
//func (this *Weixin) Login(){
//	code := this.GetString("code")
//	if code == ""{
//		models.NError(this.Ctx,"缺少code参数")
//		return
//	}
//
//	userAccToken,err := service.WeixinSvr.GetUserAccessToken(code)
//	if err != nil{
//		models.NError(this.Ctx,err.Error())
//		return
//	}
//
//	//获取用户信息
//	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN",userAccToken.AccessToken,userAccToken.Openid)
//	resp,err := http.Get(url)
//	if err != nil{
//		models.NError(this.Ctx,err.Error())
//		return
//	}
//	defer resp.Body.Close()
//	userDataByte,err := ioutil.ReadAll(resp.Body)
//	err = json.Unmarshal(userDataByte,&this.UserInfo)
//	if err != nil{
//		models.NError(this.Ctx,err.Error())
//		return
//	}
//
//
//	err = initialize.O.Raw("select id,avatar,nickname,sex from qagame_users where openid = ?",this.UserInfo.OpenID).QueryRow(&this.AuthRes)
//	if err != nil{
//		if err == orm.ErrNoRows{
//			res,err := initialize.O.Raw("insert qagame_users set avatar = ?, nickname = ?,openid = ?,sex = ?",this.UserInfo.Headimgurl,this.UserInfo.Nickname,this.UserInfo.OpenID,this.UserInfo.Sex).Exec()
//			if err != nil{
//				models.NError(this.Ctx,err.Error())
//				return
//			}
//			uid,err := res.LastInsertId()
//			if err != nil{
//				models.NError(this.Ctx,err.Error())
//				return
//			}
//			tokenStr,err := service.TokenSvr.CreateToken(uid)
//			if err != nil{
//				models.NErrorToken(this.Ctx,err.Error())
//				return
//			}
//
//			this.AuthRes.Avatar = this.UserInfo.Headimgurl
//			this.AuthRes.Id = int(uid)
//			this.AuthRes.Token = tokenStr
//			this.AuthRes.Nickname = this.UserInfo.Nickname
//			this.AuthRes.Sex = this.UserInfo.Sex
//			models.NSuccess(this.Ctx,this.AuthRes)
//			return
//		}else {
//			models.NError(this.Ctx,err.Error())
//			return
//		}
//	}
//
//	//这个用户已注册,加入更新用户信息队列
//	service.UserUpdateChan <- this.UserInfo
//
//	tokenStr,err := service.TokenSvr.CreateToken(int64(this.AuthRes.Id))
//	if err != nil{
//		models.NErrorToken(this.Ctx,err.Error())
//		return
//	}
//	this.AuthRes.Token = tokenStr
//	models.NSuccess(this.Ctx,this.AuthRes)
//	return
//}
//
//func (this *Weixin) VerifyToken() {
//	token := this.GetString("token")
//	if token == ""{
//		models.NErrorToken(this.Ctx,"缺少token")
//		return
//	}
//
//	_,err := service.TokenSvr.VerifyToken(token)
//	if err != nil{
//		models.NErrorToken(this.Ctx,err.Error())
//		return
//	}
//	models.NSuccess(this.Ctx,token)
//	return
//
//}
//
//func (this *Weixin) TestMain() {
//	redirect_uri := this.GetString("redirect_uri")
//	state := this.GetString("state")
//	this.Ctx.Redirect(301,fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx19f2c0be339ede12&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=%s#wechat_redirect",redirect_uri,state))
//	return
//}
//
