package mysql

import "log"

type Comment struct {
	Id     int    `xorm:"not null pk autoincr comment('主键') INT(11)"`
	Desc   string `xorm:"VARCHAR(255)"`
	Author int    `xorm:"INT(11)"`
}
type CommentDao struct {
	DBDao
}

func CreateCommentDao() (object *CommentDao) {
	object = new(CommentDao)
	object.Create()
	return
}
func (this *CommentDao) create(comment Comment) {
	this.InitSession()
	num, err := this.Session.Insert(comment)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("create articel num:%v", num)
	}
}
func (this *CommentDao) Comments() (comments []*Comment) {
	comments = make([]*Comment, 0)
	this.InitSession()
	this.Session.Asc("id").Find(&comments)
	return
}
func (this *CommentDao) Comment(id int) (comment *Comment) {
	comment = new(Comment)
	this.InitSession()
	this.Session.Asc("id").Where("id=? ", id).Get(comment)
	return
}
