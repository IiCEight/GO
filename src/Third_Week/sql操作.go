package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	dsn := "root:I_L0ve_yyyyy0u@tcp(127.0.0.1: 5200)/saber"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("WA")
		return
	}
	if err := db.Ping(); err != nil {
		fmt.Println("Wrong Answer!!!", err)
		return
	}
	fmt.Println("AC!")
}

func query() {
	var a int
	var b string
	var c string
	var d string
	//(1)查询服装尺码为'S'且销售价格在100以下的服装信息。
	outp := db.QueryRow("select * from costume where size = \"S\" and price < 100;")
	outp.Scan(&b, &c, &a, &d)
	fmt.Printf("code:%v, size:%v, prcie:%v, type:%v\n", b, c, a, d)
	//(2)查询仓库容量最s大的仓库。
	outp = db.QueryRow("select max(capacity) from repository;")
	outp.Scan(&a)
	fmt.Printf("maxcapacity%v\n", a)
	//（3）查询A类服装的库存总量。
	outp = db.QueryRow("select count(*) from costume where cos_type = \"A\";")
	outp.Scan(&a)
	fmt.Printf("conut:%v\n", a)
	//(4) 查询服装编码以‘A’开始开头的服装。
	roupt, err := db.Query("select * from costume where cos_encoding like \"?%\";", 'h')
	if err != nil {
		fmt.Println("WA")
	}
	for roupt.Next() {
		roupt.Scan(&b, &c, &a, &d)
		fmt.Printf("code:%v, size:%v, prcie:%v, type:%v\n", b, c, a, d)
	}
	roupt.Close()
	//（5）查询服装质量等级有不合格的供应商信息。
	roupt, err = db.Query("select * from suplier where statusofsupplier = \"?%\";", "false")
	if err != nil {
		fmt.Println("WA")
	}
	for roupt.Next() {
		roupt.Scan(&b, &c, &d)
		fmt.Printf("sucode:%v, cocode:%v, quality:%v\n", b, c, d)
	}
	roupt.Close()
}

func update() {

}

func insert() {
	// 每个表各插入一个值
	_, err := db.Exec("insert into repository value(1654, \"saber\");")
	if err != nil {
		fmt.Println("WA")
		return
	}
	_, err = db.Exec("insert into costume value(\"I\", \"Love\", 10, \"saber\");")
	if err != nil {
		fmt.Println("WA")
		return
	}
	_, err = db.Exec("insert into suplier value(\"I\", \"saber\");")
	if err != nil {
		fmt.Println("WA")
		return
	}
	_, err = db.Exec("insert into statusofsupplier value(\"I\", \"Love\", \"saber\");")
	if err != nil {
		fmt.Println("WA")
		return
	}
}

func Delete() {
	_, err := db.Exec("delete from suplier where quality = \"false\";")
	if err != nil {
		fmt.Println("WA")
		return
	}
}

func modify() {
	_, err := db.Exec("update costume set price = price * 1 + price / 10 where cos_type = \"S\";")
	if err != nil {
		fmt.Println("WA")
		return
	}

}

func main() {
	query()
	insert()
	modify()
	Delete()
}
