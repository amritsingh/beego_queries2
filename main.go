package main

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	orm.RegisterDataBase("default", "mysql", "beego_queries:tmp_pwd@tcp(127.0.0.1:3306)/beego_queries")
	orm.RegisterModel(new(User))

	o := orm.NewOrm()

	// Select one
	var user User
	err := o.QueryTable("user").Filter("id", 1).One(&user)
	fmt.Println("User with id 1: ", user, err)

	fmt.Println("------------------------------------")

	// Select one
	err = o.QueryTable("user").Filter("name", "RainbowRanger").One(&user)
	fmt.Println("User with name RainbowRanger: ", user, err)

	fmt.Println("------------------------------------")

	// Select all
	var users []User
	_, err = o.QueryTable("user").All(&users)
	fmt.Println("All users: ", users, err)

	fmt.Println("------------------------------------")

	// Select all with filter
	num, err := o.QueryTable("user").Filter("name", "RainbowRanger").All(&users)
	fmt.Println("All users with name 'RainbowRanger': ", users, num, err)

	fmt.Println("------------------------------------")

	// Select all with multiple filters
	num, err = o.QueryTable("user").Filter("name", "RainbowRanger").Filter("age__gt", 30).All(&users)
	fmt.Println("Select with Multiple filters: ", users, num, err)

	fmt.Println("------------------------------------")

	// Limit, Offset and Order
	num, err = o.QueryTable("user").Limit(10).Offset(20).OrderBy("-name").All(&users)
	fmt.Println("Select with Limit, Offset and Order: ", users, num, err)
}
