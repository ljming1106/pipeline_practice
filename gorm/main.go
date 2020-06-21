package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Account struct {
	UserID   string `gorm:"primary_key"`
	Password string
}

func initDB(db *gorm.DB) {
	db.AutoMigrate(&Account{})
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/ljm?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Panicf("db open fail!!!err:", err)
	}

	initDB(db)

	//测试事务开始，不rollback，后面是否的db操作是否可以继续
	runTransaction(db)

	db.Create(&Account{UserID: "Giraffe1"})
	fmt.Println("create done!!!")
	db.Create(&Account{UserID: "Giraffe2"})
	fmt.Println("create done1!!!")
	db.Create(&Account{UserID: "Giraffe3"})
	fmt.Println("create done2!!!")

}

func runTransaction(db *gorm.DB) (err error) {
	tx := db.Begin()
	// defer tx.Rollback()
	err = tx.Create(&Account{UserID: "Giraffe6"}).Error

	if err != nil {
		log.Println("transaction err : ", err)
		return
	}
	tx.Commit()

	// // 注意，一旦你在一个事务中，使用tx作为数据库句柄
	// if err := tx.Create(&Account{UserID: "Giraffe"}).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// if err := tx.Create(&Account{UserID: "Lion"}).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// tx.Commit()
	return
}
