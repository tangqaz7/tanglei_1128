package gin_router
import (
	api "tanglei_1128/lv1/gin_api"

	"github.com/gin-gonic/gin"
)
func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	//使用gin的Default方法创建一个路由handler
	router := gin.Default()
	//设置默认路由当访问一个错误网站时返回
	router.NoRoute(api.NotFound)
	//使用以下gin提供的Group函数为不同的API进行分组
	v1 := router.Group("admin")
	{
		v1.GET("/register", api.Register)
		v1.GET("/login", api.Login)
	}
	//监听服务器端口
	router.Run(":8080")
}
