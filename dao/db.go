package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	// TODO: encrypt and move into configuration file.
	dbType   = "mysql"
	user     = "godealuser"
	pass     = "GoDeal1234"
	protocol = "tcp"
	host     = "localhost"
	port     = "3306"
	db       = "godeal"
	charset  = "utf8"
	source   = user + ":" + pass + "@" + protocol + "(" + host + ":" + port + ")/" + db + "?charset=" + charset
)

func open() (*sql.DB, error) {
	return sql.Open(dbType, source)
}

func main() {  //main函数

	db, err := open()
	//数据库连接字符串，别告诉我看不懂。端口一定要写/
	if err != nil {  //连接成功 err一定是nil否则就是报错
		panic(err.Error()) //抛出异常
		fmt.Println(err.Error())//仅仅是显示异常
	}
	defer db.Close()  //只有在前面用了 panic 这时defer才能起作用
	db.Query("insert into demo (v) VALUES ('測試一下呀1')")
	db.Query("insert into demo (v) VALUES ('測試一下呀2')")
	rows, err := db.Query("select v from demo")
	if err != nil {
		panic(err.Error())
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var v string //定义新闻标题
	// golang的rows 好比数据库的游标，需要用scan函数把对应的值扫进去.当然也可以自己循环它的属性索引
	// 不过不建议这么做。程序可读性太差
	for rows.Next() { //开始循环、像游标吗？必须rows.Next()哦
		rerr := rows.Scan(&v) //扫每一行，并把字段的值赋到id和newstitle里面去
		if rerr == nil {
			fmt.Println(v) //输出来而已，看看
		}

	}
	db.Close() //关闭数据库 别告诉我 你不想关
}
