package main

import (
	"database/sql"
	dbu "devdungeon/go-file-tutorial/5_write_to_database_posts/dbp"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"time"
)

var db *sql.DB
var err error


func main() {
	db, err = sql.Open(dbu.DbDriver, dbu.DataSourceName)
	dbu.CheckSuccesConnection(err)
	defer db.Close()

	urlPosts := "https://jsonplaceholder.typicode.com/posts?userId=7"
	conn, err := http.Get(urlPosts)
	dbu.CheckErr(err)
	defer conn.Body.Close()

	r, err := ioutil.ReadAll(conn.Body)
	dbu.CheckErr(err)

	go insertPostInfo(&r)

	time.Sleep(time.Millisecond * 360)
}

func insertPostInfo(r *[]byte) {
	var post []dbu.InfoPost
	err = json.Unmarshal(*r, &post)
	dbu.CheckErr(err)

	for i := range post {
		sqlStmt := fmt.Sprintf(" INSERT %s SET userID = ?, id = ?, title = ?, body = ? ", dbu.TablePost)
		stmt, err := db.Prepare(sqlStmt)
		dbu.CheckErr(err)

		_, err = stmt.Exec(post[i].UserId, post[i].Id, post[i].Title, post[i].Body)
		dbu.CheckErr(err)

		id := &post[i].Id

		go insertComments(id)
	}

}

func insertComments(id *int) {
	urlComments := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%v", *id )
	conn, err := http.Get(urlComments)
	dbu.CheckErr(err)

	var comment []dbu.InfoComment
	r, err := ioutil.ReadAll(conn.Body)
	dbu.CheckErr(err)

	err = json.Unmarshal(r, &comment)
	dbu.CheckErr(err)

	for i := range comment {
		sqlStmt := fmt.Sprintf("INSERT %s SET postID = ?, id = ?, name = ?, email = ?, body = ?  ", dbu.TableComments)
		stmt, err := db.Prepare(sqlStmt)
		dbu.CheckErr(err)

		_,err = stmt.Exec(comment[i].PostId, comment[i].Id, comment[i].Name, comment[i].Email, comment[i].Body)
		dbu.CheckErr(err)
	}

}
