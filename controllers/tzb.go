package controllers
//
//import (
//	"server_admin/initialize"
//	"server_admin/models"
//	"time"
//	//"MagicEnglish/common/sdk/gowechat/mch/pay"
//	"strconv"
//	"io/ioutil"
//	"encoding/xml"
//)
//
//var testM = make(map[string]string)
//var payed = false
//
//func (this *Weixin) GetJsSdk() {
//	ctx := this.Ctx
//	url := this.GetString("url")
//	if url == ""{
//		models.NError(ctx,"")
//		return
//	}
//	js := initialize.Mp.GetJs()
//	conf,err := js.GetConfig(url)
//	if err != nil{
//		models.NError(ctx,err.Error())
//		return
//	}
//	models.NSuccess(ctx,conf.ToMap())
//	return
//}
//
//func (this *Weixin) GetPay() {
//
//	if payed{
//		models.NSuccess(this.Ctx,testM)
//		initialize.Log.Debug("缓存支付=======")
//		return
//	}
//
//	amount,err := this.GetInt("amount") //元
//	if err != nil{
//		models.NError(this.Ctx,err.Error())
//		return
//	}
//	uid := this.Ctx.Input.GetData("uid")
//	openid := ""
//	initialize.O.Raw("select openid from qagame_users where id = ?",uid).QueryRow(&openid)
//	time := time.Now().Unix()
//	order_num := strconv.FormatInt(time, 10)
//	var order = pay.OrderInput {
//		OpenID: openid, //trade_type=JSAPI时（即公众号支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识
//		Body: "购物支付",       //String(128)
//		OutTradeNum: order_num, //String(32) 20150806125346 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
//		TotalFee: amount * 100,     //分为单位
//		IP: "60.15.187.88",
//		NotifyURL:  "https://momoman.cn/v1/weixin/tzb/openapi/pay/callback", //异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数
//		ProductID:"",    //trade_type=NATIVE时（即扫码支付），此参数必传
//		//tradeType: "JSAPI", //JSAPI，NATIVE，APP
//	}
//	mch,err := initialize.Wc.MchMgr()
//	if err != nil{
//		models.NError(this.Ctx,err.Error())
//		return
//	}
//	cfg, err := mch.GetPay().GetJsAPIConfig(order)
//	if err != nil{
//		models.NError(this.Ctx,err.Error())
//		return
//	}
//	mapData := cfg.ToMap()
//	testM = mapData
//	payed = true
//	models.NSuccess(this.Ctx,mapData)
//	return
//}
//
////<xml><appid><![CDATA[wx19f2c0be339ede12]]></appid>
////<bank_type><![CDATA[CFT]]></bank_type>
////<cash_fee><![CDATA[100]]></cash_fee>
////<fee_type><![CDATA[CNY]]></fee_type>
////<is_subscribe><![CDATA[N]]></is_subscribe>
////<mch_id><![CDATA[1493165212]]></mch_id>
////<nonce_str><![CDATA[35vNF]]></nonce_str>
////<openid><![CDATA[oBYAkw3URP9pAQekMZ1GYmuNfFfQ]]></openid>
////<out_trade_no><![CDATA[1554170894]]></out_trade_no>
////<result_code><![CDATA[SUCCESS]]></result_code>
////<return_code><![CDATA[SUCCESS]]></return_code>
////<sign><![CDATA[1EB04354932A7DB183016AFFFA7B049C]]></sign>
////<time_end><![CDATA[20190402100821]]></time_end>
////<total_fee>100</total_fee>
////<trade_type><![CDATA[JSAPI]]></trade_type>
////<transaction_id><![CDATA[4200000296201904028427680928]]></transaction_id>
////</xml>
//
//type WpayCallReq struct {
//	Appid string `xml:"appid"`
//	BankType string `xml:"bank_type"`
//	CashFee string `xml:"cash_fee"`
//	FeeType string `xml:"fee_type"`
//	IsSubscribe string `xml:"is_subscribe"`
//	MchId string `xml:"mch_id"`
//	NonceStr string `xml:"nonce_str"`
//	Openid string `xml:"openid"`
//	OutTradeNo string `xml:"out_trade_no"`
//	ResultCode string `xml:"result_code"`
//	ReturnCode string `xml:"return_code"`
//	Sign string `xml:"sign"`
//	TimeEnd string `xml:"time_end"`
//	TotalFee int `xml:"total_fee"`
//	TradeType string `xml:"trade_type"`
//	TransactionId string `xml:"transaction_id"`
//}
//
//type WpayCallRes struct {
//	ReturnCode string `xml:"return_code"`
//	ReturnMsg string `xml:"return_msg"`
//}
//
////<xml><return_code>SUCCESS</return_code><return_msg>OK</return_msg></xml>
//func(this *Weixin)WpayCall(){
//	initialize.Log.Debug("%s",this.Ctx.Request.Body)
//	d,err := ioutil.ReadAll(this.Ctx.Request.Body)
//	if err != nil{
//		initialize.Log.Error("%v",err.Error())
//		return
//	}
//	req := WpayCallReq{}
//	err = xml.Unmarshal(d,&req)
//	if err != nil{
//		initialize.Log.Error("解析错误:%v",err)
//		return
//	}
//	initialize.Log.Debug("返回===%+v/n",req)
//
//	res := WpayCallRes{}
//	res.ReturnCode = "SUCCESS"
//	res.ReturnMsg = "OK"
//	this.Ctx.Output.XML(res,false)
//	return
//}
//
////{
////	Appid:wx19f2c0be339ede12
////	BankType:CFT CashFee:100
////	FeeType:CNY
////	IsSubscribe:N
////	MchId:1493165212
////	NonceStr:XYR9a
////	Openid:oBYAkw3URP9pAQekMZ1GYmuNfFfQ
////	OutTradeNo:1554172624
////	ResultCode:SUCCESS
////	ReturnCode:SUCCESS
////	Sign:3F98C6E8A87CEDDB64B2040C2D5AD771
////	TimeEnd:20190402103708
////	TotalFee:100
////	TradeType:JSAPI
////	TransactionId:4200000291201904028339807508
////}
//
//
//
