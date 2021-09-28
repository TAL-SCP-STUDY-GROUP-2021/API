package mysql

import (
	"fmt"
	"log"
)

type User struct {
	Id     int    `xorm:"not null pk autoincr comment('主键') INT(11)"`
	Age    int    `xorm:"not null default 0 comment('年龄') INT(11)"`
	Sex    int    `xorm:"not null default 0 comment('性别') INT(11)"`
	Email  string `xorm:"not null default '0' comment('email') VARCHAR(50)"`
	Friend string `xorm:"not null default '0' comment('好友') VARCHAR(50)"`
	Name   string `xorm:"not null comment('姓名') VARCHAR(255)"`
}
type UserDao struct {
	DBDao
}

func CreateUserDao() (object *UserDao) {
	object = new(UserDao)
	object.Create()
	return
}
func (this *UserDao) selectUserByID(id int) ([]User, error) {
	this.InitSession()
	//this.Session.Table("user")
	users := make([]User, 0)
	fmt.Printf("%v", len(users))
	err := this.Session.Asc("id").Where("id>? ", id).Limit(1000, 0).Find(&users)
	return users, err
}
func (this *UserDao) create(user User) {
	this.InitSession()
	num, err := this.Session.Insert(user)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("create articel num:%v", num)
	}
}
func (this *UserDao) Users() (users []*User) {
	users = make([]*User, 0)
	this.InitSession()
	this.Session.Asc("id").Find(&users)
	return
}
func (this *UserDao) User(id int) (user *User) {
	user = new(User)
	this.InitSession()
	this.Session.Asc("id").Where("id=? ", id).Get(user)
	return
}
