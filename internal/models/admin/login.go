package admin

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRes struct {
	Uid      int    `json:"uid"`
	Token    string `json:"token"`
	Username string `json:"username"`
}
