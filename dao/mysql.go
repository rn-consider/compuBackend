package dao

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
	sqlDB *sql.DB
)
func InitMySQL()(err error){
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	//注意!这句不能用:=否则会报空指针错误
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil{
		return
	}
	// 获取通用数据库对象 sql.DB，然后使用其提供的功能
	//注意!这句不能用:=否则会报空指针错误
	sqlDB, err = DB.DB()
	return sqlDB.Ping()
	
}
func Close(){
	sqlDB.Close()
}