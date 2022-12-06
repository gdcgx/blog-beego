package utils

import (
	"bytes"
	"crypto/md5"
	"database/sql"
	"fmt"
	"html/template"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
)

var db *sql.DB

func InitMysql() {
	fmt.Println("InitMysql....")
	driverName := beego.AppConfig.String("driverName")

	//注册数据库驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)

	//数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/cmsproject?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	var err error
	db, err = sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		CreateTableWithUser()
		CreateTableWithArticle()
	}
}

// 操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

// 创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDB(sql)
}

func CreateTableWithArticle() {
	sql := `CREATE TABLE IF NOT EXISTS article(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		title VARCHAR(30),
		author VARCHAR(20),
		tags VARCHAR(30),
		short VARCHAR(255),
		content LONGTEXT,
		createtime INT(10)
		);`
	ModifyDB(sql)
}

// 查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

// 传入的数据不一样，那么MD5后的32位长度的数据肯定会不一样
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

/**
 * 将文章详情的内容，转换成HTMl语句
 */
func SwitchMarkdownToHtml(content string) template.HTML {

	markdown := blackfriday.MarkdownCommon([]byte(content))

	//获取到html文档
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))

	/**
	对document进程查询，选择器和css的语法一样
	第一个参数：i是查询到的第几个元素
	第二个参数：selection就是查询到的元素
	*/
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
		fmt.Println(selection.Html())
		fmt.Println("light:", string(light))
		fmt.Println("\n\n\n")
	})
	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}
