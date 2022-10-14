package routes

import (
	"gin-blog/api/v1"
	"gin-blog/middleware"
	"gin-blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRoute() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		// auth.GET("user/info:id", v1.GetUserInfo)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		//文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		//上传文件
		auth.POST("upload", v1.UpLoad)
	}
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("category", v1.GetCategorys)
		router.GET("article", v1.GetArticles)
		router.GET("article/list/:id", v1.GetCateArt)
		router.GET("article/info/:id", v1.GetArticleInfo)
		router.POST("login", v1.Login)
		router.GET("user/info:id", v1.GetUserInfo)
	}
	_ = r.Run(utils.HttpPort)
}
