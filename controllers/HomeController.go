package controllers

import "fmt"

type HomeController struct {
	//beego.Controller
	BaseController
}

func (it *HomeController) Get() {
	fmt.Println("IsLogin:", it.IsLogin, it.LoginUser)
	it.TplName = "home.html"
}
