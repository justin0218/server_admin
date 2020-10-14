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
	GetMall() *gorm.DB
}

type mysql struct {
	masterOnce sync.Once
	mallOnce   sync.Once
	masterDb   *gorm.DB
	mallDb     *gorm.DB
	err        error
}

func (this *mysql) Get() *gorm.DB {
	this.masterOnce.Do(func() {
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

func (this *mysql) GetMall() *gorm.DB {
	this.mallOnce.Do(func() {
		this.mallDb, this.err = gorm.Open("mysql", configs.Dft.Get().Mysql.Mall.Addr)
		if this.err != nil {
			panic(this.err)
		}
		if configs.Dft.Get().Runmode == "debug" {
			this.mallDb.LogMode(true)
		}
		this.mallDb.DB().SetMaxOpenConns(configs.Dft.Get().Mysql.Mall.MaxOpenConns)
		this.mallDb.DB().SetMaxIdleConns(configs.Dft.Get().Mysql.Mall.MaxIdleConns)
	})
	return this.mallDb
}
