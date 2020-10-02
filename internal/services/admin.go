package services

import (
	"github.com/parnurzeal/gorequest"
	"server_admin/api"
	"server_admin/internal/models/admin"
	"server_admin/pkg/jwt"
)

type AdminService struct {
}

func (s *AdminService) Login(username string, password string) (ret admin.LoginRes, err error) {
	user := new(admin.User)
	loginRes, e := user.Login(username, password)
	if e != nil {
		err = e
		return
	}
	ret.Username = loginRes.Username
	ret.Token, err = jwt.CreateToken(int64(loginRes.Uid))
	ret.Uid = loginRes.Uid
	return
}

func (s *AdminService) FileRead(key string) (ret string, err error) {
	client := api.Rds.Get()
	res, err := client.Get(key).Result()
	if err == nil && res != "" {
		ret = res
		return
	}
	_, retBytes, errs := gorequest.New().Get(key).End()
	if len(errs) > 1 {
		err = errs[0]
		return
	}
	ret = string(retBytes)
	err = client.Set(key, ret, -1).Err()
	return
}
