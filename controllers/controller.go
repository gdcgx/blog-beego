package controllers

import (
	"fmt"
	"testBeego/models"
	"testBeego/utils"
	"time"

	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (t *RegisterController) Get() {
	t.TplName = "register.html"
}

// 处理注册
func (t *RegisterController) Post() {
	//获取表单信息
	username := t.GetString("username")
	password := t.GetString("password")
	repassword := t.GetString("repassword")
	if password != repassword {
		t.Data["json"] = map[string]interface{}{"code": 0, "message": "两次输入的密码不一致"}
		t.ServeJSON()
		return
	}
	fmt.Println(username, password, repassword)

	//注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	id := models.QueryUserWithUsername(username)
	fmt.Println("id:", id)
	if id > 0 {
		t.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已经存在"}
		t.ServeJSON()
		return
	}

	//注册用户名和密码
	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	password = utils.MD5(password)
	fmt.Println("md5后:", password)

	user := models.User{
		Id:         0,
		Username:   username,
		Password:   password,
		Status:     0,
		Createtime: time.Now().Unix(),
	}
	_, err := models.InsertUser(user)
	if err != nil {
		t.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		t.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}
	t.ServeJSON()
}
