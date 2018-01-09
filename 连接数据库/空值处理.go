package  main

import ( "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	// 创建数据库对象
	db, err := sql.Open("mysql", "qiang:qiang@tcp(10.9.184.249:3307)/qiang?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 现在数据库有一张表qiang.test1  数据如下

	/*  mysql> select * from test1;
		+------+------+------+------+
		| a    | b    | c    | d    |
	    +------+------+------+------+
		|  1 | NULL |    2 | age  |
	    +------+------+------+------+
	*/

	rows,err := db.Query("SELECT * FROM qiang.test1")
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()

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
}