package  main

import ( "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	//创建数据库对象
	db, err := sql.Open("mysql","qiang:qiang@tcp(10.9.184.249:3307)/qiang?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//执行数据库操作  创建一张表
	_,err = db.Exec("CREATE TABLE IF NOT EXISTS qiang.hello(world varchar(50))")
	if err != nil{
		log.Fatal(err)
	}

	//插入数据
	rs, err := db.Exec("INSERT INTO qiang.hello(world) VALUES('hello world')")
	if err != nil{
		log.Fatal(err)
	}
	rowCount, err := rs.RowsAffected()
	if err != nil{
		log.Fatal(err)
	}
	log.Printf("inserted %d rows",rowCount)
	//返回结果 2018/01/08 18:42:10 inserted 1 rows

	//查询数据
	rows, err := db.Query("SELECT  world FROM qiang.hello")
	if err != nil{
		log.Fatal(err)
	}
	//query方法执行select查询语句 返回的是一个sql.Rows类型的结果集 迭代后者的next的方法 然后使用scan方法给变量s赋值

	for rows.Next(){
		var s string
		err := rows.Scan(&s)
		if err != nil{
			log.Fatal(err)
		}
		log.Printf("found row containing %q",s)
	}
	rows.Close()

	// 连接池:当你的函数(例如Exec，Query)调用需要访问底层数据库的时候，函数首先会向连接池请求一个连接。如果连接池有空闲的连接
	// 则返回给函数。否则连接池将会创建一个新的连接给函数。一旦连接给了函数，连接则归属于函数。函数执行完毕后，要不把连接所属权
	// 归还给连接池，要么传递给下一个需要连接的（Rows）对象，最后使用完连接的对象也会把连接释放回到连接池。

	// 请求一个连接的函数有好几种，分以下几种：
	// db.Ping() 调用完毕后会马上把连接返回给连接池
	// db.Exec() 调用完毕后会马上把连接返回给连接池，但是它返回的Result对象还保留这连接的引用，当后面的代码需要处理结果集的时候连接将会被重用。
	// db.Query() 调用完毕后会将连接传递给sql.Rows类型，当然后者迭代完毕或者显示的调用.Close()方法后，连接将会被释放回到连接池。
	// db.QueryRow()调用完毕后会将连接传递给sql.Row类型，当.Scan()方法调用之后把连接释放回到连接池。
	// db.Begin() 调用完毕后将连接传递给sql.Tx类型对象，当.Commit()或.Rollback()方法调用后释放连接。

	//通过db.Ping()方法初始化
	err = db.Ping()
	if err != nil{
		log.Fatal(err)
	}

	// 连接池配置
	// db.SetMaxOpenConns(n int) 设置打开数据库的最大连接数 包含正在使用的连接和连接池的连接 如果你的函数需要申请一个连接，并且连接池已经
	// 没有了新的连接或者连接数达到了最大连接数 此时的函数调用将会被block 知道有可用的连接才会返回
	// 设置这个值是为了避免并发太高导致MySQL出现too many connections的错误 该函数的默认值是0  表示无限制

	// db.SetMaxIdleConns(n int) 设置连接池中的保持连接的最大连接数，默认也是0，表示当连接释放回连接池的时候，连接将会被关闭，这会导致连接
	// 池中频繁的关闭和创建



	//读取数据 1.从连接池中请求一个连接 2.执行查询的sql语句 3.将数据库连接的所属权传递给Result结果集
	rows,err = db.Query("SELECT world FROM qiang.hello")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("foun row containing %q",s)
	}
	rows.Close()





}
