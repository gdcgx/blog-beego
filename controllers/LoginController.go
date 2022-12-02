package controllers

import (
	"fmt"
	"testBeego/models"
	"testBeego/utils"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (it *LoginController) Get() {
	it.TplName = "login.html"
}

func (it *LoginController) Post() {
	username := it.GetString("username")
	password := it.GetString("password")
	fmt.Println("username:", username, ",password", password)

	id := models.QueryUserWithParam(username, utils.MD5(password))
	fmt.Println("id:", id)
	if id > 0 {
		// 设置了session后悔数据处理设置到cookie，然后再浏览器进行网络请求的时候回自动带上cookie
		// 因为我们可以通过获取这个cookie来判断用户是谁，这里我们使用的是session的方式进行设置
		it.SetSession("loginuser", username)
		it.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		it.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	it.ServeJSON()
}
