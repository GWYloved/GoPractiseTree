package main

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

func main()  {
	db, err := sql.Open("sqlite3","./foo.db")
	checkErr(err)

	sql_table := `
	CREATE TABLE IF NOT EXISTS userinfo(
		uid INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(64) NULL,
		departname VARCHAR(64) NULL,
		created DATE NULL
	);`

	db.Exec(sql_table)

	//insert
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values (?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("ypc","haomaiyi","2018-03-07")
	checkErr(err)

	_, err = stmt.Exec("ccm","chuanyin","2018-01-01")
	checkErr(err)

	id, err := res.LastInsertId() //返回自增数字
	checkErr(err)

	fmt.Println("id = "+string(id))

	//update
	stmt, err = db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec("ypc_new",id)
	checkErr(err)

	affect, err := res.RowsAffected() //返回的是此次操作影响的行数
	checkErr(err)

	fmt.Printf("rows affect = "+ string(affect))

	//query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next(){
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid = "+ string(uid))
		fmt.Println("username = "+username)
		fmt.Println("department = "+department)
		fmt.Println("create = "+created.String())
	}
	rows.Close()

	//delete
	stmt, err = db.Prepare("delete FROM userinfo WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println("affect = "+string(affect))

	db.Close()

}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}