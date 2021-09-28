package mysql

import (
	"Graphql/app/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

type DBDao struct {
	Engine  *xorm.Engine
	Session *xorm.Session
}

func (this *DBDao) Create() {
	var err error
	this.Engine, err = xorm.NewEngine("mysql", config.Db)
	if err != nil {
		log.Println(err)
	}
	this.Engine.ShowSQL(true)
	return
}
func (this *DBDao) InitSession() {
	this.Session = this.Engine.NewSession()

}

func (this *DBDao) Close() {
	this.Engine.Close()
}
