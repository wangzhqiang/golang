package  main

import ( "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

func main() {

	// 创建数据库对象
	db, err := sql.Open("mysql", "qiang:qiang@tcp(10.9.184.249:3307)/qiang?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM qiang.test1")
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()  //记得关闭连接

	cols, err := rows.Columns() //rows.Columns得到的是col作为map的key值
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(cols)         //返回值为 [a b c d]


	vals := make([][]byte,len(cols))     // [[] [] [] []]长度为4的slice
	scans := make([]interface{},len(cols)) //[<nil> <nil> <nil> <nil>]长度为4的slice

	for i := range vals{
		scans[i] = &vals[i]    //将vals的地址赋值给scans [0xc4200a6060 0xc4200a6078 0xc4200a6090 0xc4200a60a8]
	}

	var results []map[string]string

	for rows.Next(){
		err = rows.Scan(scans...) //其实是将之前定义的vals的地址传给Scan方法 这时候vals=[[49] [] [50] [97 103 101]]
		if err != nil{
			log.Fatal(err)
		}

		row := make(map[string]string)
		for k,v := range vals{       //循环取值赋值
			key := cols[k]
			row[key] = string(v)
		}
		results = append(results,row)
	}

	for k,v := range results{
		fmt.Println(k,v)
	}
}
