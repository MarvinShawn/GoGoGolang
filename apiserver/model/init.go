package model

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Self *gorm.DB
	Docker *gorm.DB
}

var DB *Database

func openDB(username,password,addr,name string) *gorm.DB  {

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")

	db,err := gorm.Open("mysql",config)

	if err != nil {
		log.Errorf(err,"Database connection failed. Database name: %s",name)
	}

	setupDB(db)

	return db

}

func setupDB(db *gorm.DB)  {

	db.LogMode(viper.GetBool("gormlog"))
	//db.DB().SetMaxOpenConns(20000)  用于设置最大打开的连接数，默认值为0表示不限制。设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections
	db.DB().SetMaxIdleConns(0) //用于设置闲置的连接数，设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用

}


func InitSelfDB() *gorm.DB  {

	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))

}

func GetSelfDB() *gorm.DB {

	return InitSelfDB()

}

func InitDockerDB() *gorm.DB  {

	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))

}

func GetDockerDB() *gorm.DB  {

	return InitDockerDB()
}

//初始化连接
func (db *Database) Init()  {

	DB = &Database{
		Self:GetSelfDB(),
		Docker:GetDockerDB(),
	}

}

func (db *Database) Close()  {

	DB.Self.Close()
	DB.Docker.Close()

}







