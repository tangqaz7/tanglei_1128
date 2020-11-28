package gin_regAndlog
import (
	"net/http"
	. "tanglei_1128/lv1/data"
	Func "tanglei_1128/lv1/func_judge"

	"github.com/gin-gonic/gin"
)
//注册
func Register(c *gin.Context) {
	//获取用户名、密码
	name := c.Request.FormValue("Name")
	passwd := c.Request.FormValue("Passwd")
	//判断用户是否存在
	//存在输出状态1
	//不存在创建用户，保存密码与用户名
	Bool := Func.IsExist(name)
	if Bool {
		//注册状态
		State["state"] = 1
		State["text"] = "此用户已存在！"
	} else {
		//用户不存在即添加用户
		AddStruct(name, passwd)
		State["state"] = 1
		State["text"] = "注册成功！"
	}

	//把状态码和注册状态返回给客户端
	c.String(http.StatusOK, "%v", State)
}
//登录
func Login(c *gin.Context) {
	name := c.Request.FormValue("Name")
	passwd := c.Request.FormValue("Passwd")
	//先判断用户是否存在，存在再判断密码是否正确
	Bool := Func.IsExist(name)
	if Bool {
		Bool_Pwd := Func.IsRight(name, passwd)
		if Bool_Pwd {
			State["state"] = 1
			State["text"] = "登录成功！"
		} else {
			State["state"] = 0
			State["text"] = "密码错误！"
		}
	} else {
		State["state"] = 2
		State["text"] = "登录失败！此用户尚未注册！"
	}

	c.String(http.StatusOK, "%v", State)
}
//设置默认路由当访问一个错误网站时返回
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404 ,page not exists!",
	})
}
//添加用户
func AddStruct(name string, passwd string) {
	var user User
	user.Name = name
	user.Passwd = passwd
	user.Id = len(Slice) + 1
	Slice = append(Slice, user)
}