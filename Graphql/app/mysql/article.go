package mysql

import "log"

type Article struct {
	Id          int    `xorm:"not null pk autoincr comment('主键') INT(11)"`
	Text        string `xorm:"VARCHAR(255)"`
	Ispublished int    `xorm:"not null default 0 INT(11)"`
	Author      int    `xorm:"INT(11)"`
	Comments    string `xorm:"VARCHAR(255)"`
}
type ArticleDao struct {
	DBDao
}

func CreateArticleDao() (object *ArticleDao) {
	object = new(ArticleDao)
	object.Create()
	return
}

func (this *ArticleDao) create(article Article) {
	this.InitSession()
	num, err := this.Session.Insert(article)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("create articel num:%v", num)
	}
}
func (this *ArticleDao) Articles() (articles []*Article) {
	this.InitSession()
	this.Session.Asc("id").Find(&articles)
	return
}
func (this *ArticleDao) Article(id int) (article *Article) {
	article = new(Article)
	this.InitSession()
	this.Session.Asc("id").Where("id=? ", id).Get(article)
	return
}
func (this *ArticleDao) CreateArticle(article *Article) {
	this.InitSession()
	_, err := this.Session.Insert(article)
	if err != nil {

	}
	return
}
func (this *ArticleDao) LastArticle() (article *Article) {
	article = new(Article)
	this.InitSession()
	this.Session.Desc("id").Limit(1, 0).Get(article)

	return
}
