package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin   bool
	LoginUser interface{}
}

func (it *BaseController) Get() {
	it.Data["Website"] = "beego.me"
	it.Data["Email"] = "astaxie@gmail.com"
	it.TplName = "index.tpl"
}

// 判断是否登录
func (it *BaseController) Prepare() {
	loginuser := it.GetSession("loginuser")
	fmt.Println("loginuser---->", loginuser)
	if loginuser != nil {
		it.IsLogin = true
		it.LoginUser = loginuser
	} else {
		it.IsLogin = false
	}
	it.Data["IsLogin"] = it.IsLogin
}
