package main

//https://github.com/microsoft/vscode-go
import (
	"com/sqrt"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func add(a, b int) (int, string) {

	if a < 0 {
		return a + b, "ok"
	}
	return b - a, "ok"
}
func checkErr(err error) {
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
}
func main() {
	db, err := sql.Open("mysql", "root:root123@/test?charset=utf8")
	defer db.Close()
	checkErr(err)
	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("11111", "研发部门", "2012-12-09")
	id, err := res.LastInsertId()
	fmt.Println(id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	checkErr(err)
	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	//删除数据
	//	stmt, err = db.Prepare("delete from userinfo where uid=?")
	//	checkErr(err)
	//	res, err = stmt.Exec(id)
	//	checkErr(err)
	//	affect, err = res.RowsAffected()
	//	checkErr(err)
	fmt.Println(affect)
	fmt.Printf("你好 %v\n", sqrt.Sqrt(2))
	http.HandleFunc("/", sayHellowName)
	http.HandleFunc("/login/", loginHandle)
	err = http.ListenAndServe(":1111", nil)

	if err != nil {
		log.Fatal("服务器错误:", err)
	}

}
