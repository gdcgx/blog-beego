package controllers

import (
	"fmt"
	"testBeego/models"
	"time"
)

type AddArticleController struct {
	BaseController
}

/*
当访问/add路径的时候回触发AddArticleController的Get方法
响应的页面是通过TpName
*/
func (it *AddArticleController) Get() {
	it.TplName = "write_article.html"
}

func (it *AddArticleController) Post() {
	title := it.GetString("title")
	tags := it.GetString("tags")
	short := it.GetString("short")
	author := fmt.Sprintf("%v", it.GetSession("loginuser"))
	content := it.GetString("content")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	//实例化model，将它出入到数据库中
	art := models.Article{
		Id:         0,
		Title:      title,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Author:     author,
		Createtime: time.Now().Unix(),
	}
	_, err := models.AddArticle(art)

	//返回数据给浏览器
	var response map[string]interface{}
	if err == nil {
		//无误
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}

	it.Data["json"] = response
	it.ServeJSON()
}
