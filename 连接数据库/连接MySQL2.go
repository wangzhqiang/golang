package  main

import ( "database/sql"
          _ "github.com/go-sql-driver/mysql"
         "log"
 )

func main() {

	//创建数据库对象
	db, err := sql.Open("mysql", "qiang:qiang@tcp(10.9.184.249:3307)/qiang?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//读取数据 1.从连接池中请求一个连接 2.执行查询的sql语句 3.将数据库连接的所属权传递给Result结果集
	rows,err := db.Query("SELECT world FROM qiang.hello")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {        //rows.Next方法是用来迭代，当它迭代到最后一行，会触发一个io.EOF的信号，即引发错误，同时go会调用rows.Close方法释放连接
		var s string
		err = rows.Scan(&s)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("foun row containing %q",s)
	}
	rows.Close()   //如果没有关闭rows连接 将导致大连的连接并且不会被其他函数重用 最终将导致数据库无法正常使用


	//读取数据修改
	rows,err = db.Query("SELECT world FROM qiang.hello")
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil{
			log.Fatal(err)
		}
		log.Printf("found row containing %q",s)
	}
	rows.Close()
	if err = rows.Err(); err!=nil {    //为了检查是否迭代正常退出还是异常退出 需要检查rows.Err
		log.Fatal(err)
	}
}

