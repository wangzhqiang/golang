package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

// 事务用的是sql.Tx对象
// 事务的连接生命周期从Beigin函数调用起，直到Commit和Rollback函数的调用结束


func doSomething(){
	panic("A Panic Running Error")
}

func clearTransction(tx *sql.Tx){
	err := tx.Rollback()
	if err != sql.ErrTxDone && err != nil {
		log.Fatal(err)
	}
}


func main(){
	// 创建数据库对象
	db, err := sql.Open("mysql", "qiang:qiang@tcp(10.9.184.249:3307)/qiang?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

    tx,err := db.Begin() //开启一个事务
    if err != nil{
    	log.Fatal(err)
	}

	defer clearTransction(tx)

	rs, err := tx.Exec("UPDATE qiang.test1 set b ='qiang' where a=1")  //更新数据
	if err != nil{
		log.Fatal(err)
	}

	rowAffected,err := rs.RowsAffected()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rowAffected)

	rs, err = tx.Exec("UPDATE qiang.test1 set b ='text' WHERE a=2")    //更新数据
	if err != nil {
		log.Fatal(err)
	}

	rowAffected,err = rs.RowsAffected()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rowAffected)

	doSomething()

	if err := tx.Commit();err != nil{
		log.Fatal(err)
	}

}