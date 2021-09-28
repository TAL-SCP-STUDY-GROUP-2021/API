package dao

type User struct {
	Id     int    `xorm:"not null pk autoincr comment('主键') INT(11)"`
	Age    int    `xorm:"not null default 0 comment('年龄') INT(11)"`
	Sex    int    `xorm:"not null default 0 comment('性别') INT(11)"`
	Email  string `xorm:"not null default '0' comment('email') VARCHAR(50)"`
	Friend string `xorm:"not null default '0' comment('好友') VARCHAR(50)"`
	Name   string `xorm:"not null comment('姓名') VARCHAR(255)"`
}
