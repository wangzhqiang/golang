package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/* 查询
  1 rows, err := db.Query("SELECT * FROM qiang.test1 where a = 1")
  2 rows, err := db.Query("SELECT * FROM qiang.test1 where a = ?",1)
*/

 /* prepared 即带有占位符的sql语句,客户端将改语句和参数发给MySQL服务器  MySQL服务器编译成一个prepared语句,这个语句可以根据
    不同的参数多次调用

    prepared的执行方式如下
    1. 准备prepare语句
    2. 执行prepared语句和参数
    3. 关闭prepared语句

    优点:
    1.避免通过引号组装拼接SQL语句 避免SQL注入带来的风险
    2.多次执行SQL语句
  */

  func main(){
	  // 创建数据库对象
	  db, err := sql.Open("mysql", "qiang:qiang@tcp(10.9.184.249:3307)/qiang?parseTime=true")
	  if err != nil {
		  log.Fatal(err)
	  }
	  defer db.Close()

	  //自定义prepare查询
	  stmt,err := db.Prepare("SELECT * FROM qiang.test1 WHERE a = ?") //通过prepare方法创建一个stmt对象 然后执行stmt
	  if err != nil{                                                        //对象的Query方法得到sql.Rows的结果集，最后关闭stmt.Close
	  	log.Fatal(err)
	  }
	  defer stmt.Close()

	  rows, err := stmt.Query(1)
	  if err != nil {
	  	log.Fatal(err)
	  }
	  for rows.Next(){
		  var (
			  a1 int32
			  b1 sql.NullString
			  c1 int32
			  d1 string
		  )
		  err = rows.Scan(&a1,&b1,&c1,&d1)
		  if err != nil{
			  log.Fatal(err)
		  }
		  log.Printf("%q,%#v,%q,%q",a1,b1,c1,d1)
	  }
	  /*
	  单次查询不需要使用prepared，每次使用stmt语句都是三次网络请求次数，prepared execute close

      不要循环中创建prepare语句

	  注意关闭 stmt
	   */


  }