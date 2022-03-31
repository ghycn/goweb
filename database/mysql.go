package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Person struct {
	gorm.Model
	Id        int `gorm:"primary_key"`
	FirstName string
	LastName  string
}

//func init() {
//	//GetDb()
//	var err error
//	db, err = gorm.Open("mysql", "root:12345678@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
//
//
//	db.DB().SetMaxIdleConns(10)
//	db.DB().SetMaxOpenConns(100)
//	//SqlDB.SingularTable(true)
//	db.LogMode(true)
//	//db.AutoMigrate(&Person{})
//	//Scan, 原生查询
//	var person Person
//	db.Raw("SELECT id, first_name, last_name FROM person where id=?",3).Row().Scan(person)
//	fmt.Println("Scan: ", person)
//
//
//
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//}

//参数含义:数据库用户名、密码、主机ip、连接的数据库、端口号
func dbConn(User, Password, Host, Db string, Port int) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", User, Password, Host, Port, Db)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		return nil
	}
	db.SingularTable(true) //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	db.LogMode(true)       //打印sql语句
	//开启连接池
	db.DB().SetMaxIdleConns(100)   //最大空闲连接
	db.DB().SetMaxOpenConns(10000) //最大连接数
	db.DB().SetConnMaxLifetime(30) //最大生存时间(s)

	return db
}

func GetDb() (conn *gorm.DB) {
	for {
		conn = dbConn("root", "12345678", "127.0.0.1", "test", 3306)
		if conn != nil {
			break
		}
		fmt.Println("本次未获取到mysql连接")
	}
	return conn
}

type User struct {
	Id     int    `gorm:"primary_key" json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"` //1:男、2:女
}

//添加数据
func (user *User) Add() {

	conn := GetDb()
	defer conn.Close()

	err := conn.Create(user).Error
	if err != nil {
		fmt.Println("创建失败")
	}
}

//修改数据
func (user *User) Update() {
	conn := GetDb()
	defer conn.Close()

	err := conn.Model(user).Update(user).Error
	if err != nil {
		fmt.Println("修改失败")
	}
}

//删除数据
func (user *User) Del() {
	conn := GetDb()
	defer conn.Close()

	err := conn.Delete(user).Error
	if err != nil {
		fmt.Println("删除失败")
	}
}
