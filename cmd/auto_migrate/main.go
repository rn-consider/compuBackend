package main

import (
	"database/sql"
	"github.com/rn-consider/compuBackend/cmd/initialize"
	"github.com/rn-consider/compuBackend/config"
	"go.uber.org/zap"
)

func main() {
	config.GVA_VP = initialize.Viper()
	config.GVA_DB = initialize.Gorm() // gorm连接数据库
	if config.GVA_DB != nil {
		initialize.RegisterTables(config.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := config.GVA_DB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				zap.L().Error("数据表迁徙失败", zap.Error(err)) // 使用 Zap 打印数据库关闭失败的错误信息
			}
		}(db)
	}
}
