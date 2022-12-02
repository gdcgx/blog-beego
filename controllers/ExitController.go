package controllers

type ExitController struct {
	BaseController
}

func (it *ExitController) Get() {
	//清除该用户登录状态的数据
	it.DelSession("loginuser")
	it.Redirect("/", 302)
}
