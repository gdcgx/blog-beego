package main

import (
	_ "testBeego/routers"
	"testBeego/utils"

	"github.com/astaxie/beego"
)

func main() {
	utils.InitMysql()
	beego.Run()
}
