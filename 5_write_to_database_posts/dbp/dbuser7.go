package dbp

import (
	"fmt"
	"log"
)

type InfoPost struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

type InfoComment struct {
	PostId int
	Id     int
	Name   string
	Email  string
	Body   string
}

var err error

const DbDriver = "mysql"
const User = "root"
const Password = "passwordIK"
const DbName = "dbuser7"
const TablePost = "post_user7"
const TableComments = "comments_u7"

// dataSourceName := "root:passwordIK@tcp(127.0.0.1:3306)/dbuser7?charset=utf8"
var DataSourceName = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8",
	User, Password, DbName)



func CheckSuccesConnection(err error) {
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Success!")
	}

}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
