package main

import (
	json2 "encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//用户
type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Post struct {
	Id        int
	Title     string
	Content   string
	Author    string `sql:"not null"`
	CreatedAt time.Time
	Comments  []Comment
}

type Person struct {
	//gorm.Model
	//Id        int `gorm:"primary_key"`
	//FirstName string
	//LastName  string
	Name string
	Age  int
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int    `sql:"index"`
	CreatedAt time.Time
}

var DbConn *gorm.DB

func init() {
	var err error
	DbConn, err = gorm.Open("mysql", "root:12345678@/test?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
	DbConn.SingularTable(true) //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	DbConn.LogMode(true)       //打印sql语句
	//开启连接池
	DbConn.DB().SetMaxIdleConns(100)   //最大空闲连接
	DbConn.DB().SetMaxOpenConns(10000) //最大连接数
	DbConn.DB().SetConnMaxLifetime(30) //最大生存时间(s)
	DbConn.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Title: "GORM 示例教程", Content: "基于 GORM 进行数据库增删改查", Author: "学院君"}

	// 通过 GORM 插入文章记录
	DbConn.Create(&post)
	//fmt.Println(post)

	// 通过关联关系新增评论并将其附加到对应的文章记录
	comment := Comment{Content: "Test Comment", Author: "学院君小号"}
	DbConn.Model(&post).Association("Comments").Append(comment)

	// 查询文章记录
	var gormPost Post
	DbConn.Where("author = ?", "学院君").First(&gormPost)

	// 查询包含评论数据的文章记录
	//var comments []Comment
	//DbConn.Model(&gormPost).Related(&comments)
	//fmt.Println(comments[0])

	var comments []Comment
	//var comments = make([]Comment, 0)
	rows := DbConn.Raw("SELECT * FROM comments ").Scan(&comments)
	defer rows.Close()
	fmt.Println("Comment", comments)
	//for i, c := range comments {
	//	var marshal,_ = json2.Marshal(c)
	//	fmt.Println(i, string(marshal))
	//}

	var jsonArr, _ = json2.Marshal(comments)

	//var arr []Comment

	var jsonObj []interface{}
	err := json2.Unmarshal(jsonArr, &jsonObj)
	println(err, jsonObj)

	// 根据结构体生成json
	//manJson := Person{
	//	Name: "Elinx",
	//	Age:  26,
	//}

	// 解析json数组到切片（数组）
	jsonArrStr := `[{"Name":"Elinx","Age":26}, {"Name":"Twinkle","Age":21}]`
	var jsonSlice []Person

	json2.Unmarshal([]byte(jsonArrStr), &jsonSlice)

	strs := make([]string, 0)
	for i := range jsonSlice {
		marshal, _ := json2.Marshal(jsonSlice[i])
		//fmt.Println(string(marshal))
		strs = append(strs, string(marshal))
	}
	println(strs[0], strs[1])

	//stars, _ := json.Json2String(comments)
	//fmt.Println("======================")
	//println(stars)
	//
	//str := json.String2Json(stars, Comment{})
	//println(str)
	//for i := range comments {
	//	println(comments[i].Content)
	//	}
	//}

	//for rows.Next() {
	//	var comment Comment
	//	rows.Scan(&comment)
	//	comments = append(comments, comment)
	//}
	//println(comments)
	//if err = rows.Err(); err != nil {
	//	return
	//}else{
	//	println(err)
	//}
	user := User{
		UserName: "itbsl",
		NickName: "jack",
		Age:      18,
		Birthday: "2001-08-15",
		Sex:      "itbsl@gmail.com",
		Phone:    "176XXXX6688",
	}
	data, err := json2.Marshal(user)
	fmt.Printf("%s\n", string(data))

	personJson := `[{"Name":"Elinx","Age":26}, {"Name":"Twinkle","Age":21}]`
	//var user1 User
	var person []Person
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json2.Unmarshal([]byte(personJson), &person)
	fmt.Printf("%s\n", person[1].Name)

}
