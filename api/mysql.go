package api

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"server_admin/configs"
	"sync"
)

var Mysql mysqler

func init() {
	Mysql = &mysql{}
}

type mysqler interface {
	Get() *gorm.DB
}

type mysql struct {
	sync.Once
	masterDb *gorm.DB
	err      error
}

func (this *mysql) Get() *gorm.DB {
	this.Do(func() {
		this.masterDb, this.err = gorm.Open("mysql", configs.Dft.Get().Mysql.Master.Addr)
		if this.err != nil {
			panic(this.err)
		}
		if configs.Dft.Get().Runmode == "debug" {
			this.masterDb.LogMode(true)
		}
		this.masterDb.DB().SetMaxOpenConns(configs.Dft.Get().Mysql.Master.MaxOpenConns)
		this.masterDb.DB().SetMaxIdleConns(configs.Dft.Get().Mysql.Master.MaxIdleConns)
	})
	return this.masterDb
}
