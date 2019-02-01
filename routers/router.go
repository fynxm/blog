package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//首页
	beego.Router("/", &controllers.HomeController{})
	//注册页面
	beego.Router("/reg", &controllers.RegController{})
	//登录页面
	beego.Router("/login", &controllers.LoginController{})
	//退出页面
	beego.Router("/exit", &controllers.ExitController{})
	//添加文章页面
	beego.Router("/article/add", &controllers.AddArticleController{})
	//显示文章
	beego.Router("/article/:id", &controllers.ShowArticleController{})
	//修改文章
	beego.Router("/article/update", &controllers.UpdateArticleController{})
	// 删除文章
	beego.Router("/article/delete", &controllers.DeleteArticleController{})
	//标签
	beego.Router("/tags", &controllers.TagsController{})
	//相册
	beego.Router("/album", &controllers.AlbumController{})
	//文件上传
	beego.Router("/upload", &controllers.UploadController{})
	//关于我
	beego.Router("/aboutme", &controllers.AboutMeController{})
}
