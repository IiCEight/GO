package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type Resp struct {
	Encoding string
	Cap      int
}

type Cos struct {
	Encoding string
	Size     string
	Price    int
	Typ      string
}

type Sup struct {
	Encoding string
	Name     string
}

type Sta struct {
	Re_encoding string
	Co_encoding string
	Qua         string
}

func main() {
	db, err := gorm.Open("mysql", "root:439956461@(127.0.0.1:3306)/saber")
	checkerr(err)
	defer db.Close()

	//create
	// db.AutoMigrate(&Cos{}, &Resp{}, &Sta{}, &Sup{})

	// Cos1 := Cos{"saber", "S", 1345835, "sword"}
	// Into := Resp{"isaber", 20394}
	// db.Create(&Into)
	// db.Delete(&Cos1)

	var res Cos

	db.Debug().Where("Size = ? and price > ?", "S", 100).Find(&res)

	fmt.Println(res.Encoding)

	var res1 Resp

	db.Debug().Order("Cap desc").Limit(1).Find(&res1)

	fmt.Println(res1.Encoding)
	var all int
	db.Debug().Table("Cos").Where("Size = ?", "S").Count(&all)

	fmt.Println(all)

	db.Debug().Where("Encoding like ?", "S%").Find(&res)
	fmt.Println(res.Encoding)

	// var res2 Sup

	// db.Debug().Where("Qua = ?", "false").Find(&res2)

	// fmt.Println(res2.Encoding)

	res.Price = res.Price + res.Price/10

	db.Debug().Model(&res).Update("Price")

	fmt.Println(res)

	db.Delete(&res)

	return
}
