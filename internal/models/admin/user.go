package admin

import (
	"server_admin/api"
	"time"
)

const BLOG_ADMIN_USERS = "blog_admin_users"

type User struct {
	Uid int `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *User) Login(userName,passWord string) (ret User,err error)  {
	db := api.Mysql.Get().Table(BLOG_ADMIN_USERS)
	err = db.Where("username = ? and password = password(?)",userName,passWord).First(&ret).Error
	return
}
