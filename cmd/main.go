package main

import (
	"fmt"
	"server_admin/api"
	"server_admin/configs"
	"server_admin/internal/routers"
	//"server_admin/job"
)

func main() {
	api.Log.Get().Debug("starting...")
	err := api.Rds.Get().Ping().Err()
	if err != nil {
		panic(err)
	}
	api.Mysql.Get()
	fmt.Println("server started!!!")
	err = routers.Init().Run(fmt.Sprintf(":%d", configs.Dft.Get().Http.Port))
	if err != nil {
		panic(err)
	}
}
