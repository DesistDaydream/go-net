package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	//"time"
)

func main() {
	db, err := sql.Open("mysql", "root:mypassword@/caredaily?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("lichenhao", "developent", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("lichenhaoupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// 执行SQL语句查询数据，并返回查询数据
	rows, err := db.Query("SELECT * FROM type")
	checkErr(err)

	// fmt.Println(rows)

	// 迭代查询数据
	for rows.Next() {
		var id string
		var type1 string
		// var created string
		// 扫描数据中每一行的值，给每一列的数据指定一个变量
		// err = rows.Scan(&uid, &username, &department, &created)
		err = rows.Scan(&id, &type1)
		checkErr(err)
		fmt.Println("id为:", id)
		fmt.Println("type为:", type1)
		// fmt.Println("department为:", User)
		// fmt.Println("created为", created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
