package routers

import (
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"server_admin/configs"
	"server_admin/internal/controllers"
	//"server_admin/internal/middleware"
	//
	//"server_admin/internal/routers/v1/ws"
)

func Init() *gin.Engine {
	r := gin.Default()
	gin.SetMode(configs.Dft.Get().Runmode)
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "*",
		ExposedHeaders:  "",
		Credentials:     true,
		ValidateHeaders: false,
	}))
	r.GET("health", func(context *gin.Context) {
		context.JSON(200, map[string]string{"msg": "ok"})
		return
	})
	admin := new(controllers.AdminController)
	r.POST("/v1/admin/login", admin.Login)
	r.POST("/v1/open/auth/upload", admin.UploadFile)
	r.GET("/v1/admin/open/file/read", admin.FileRead)

	blog := new(controllers.BlogController)
	r.POST("/v1/admin/auth/blog/create", blog.CreateBlog)
	r.GET("/v1/admin/auth/blog/list", blog.GetBlogList)
	r.GET("v1/admin/auth/blog/detail", blog.Detail)

	bill := new(controllers.BillController)
	r.POST("/v1/admin/auth/account/bill/make", bill.Create)
	r.GET("/v1/admin/auth/account/bill/list", bill.SumBill)
	return r
}
