package initialize

import (
	"github.com/rn-consider/compuBackend/config"
	"github.com/rn-consider/compuBackend/models"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	switch config.GVA_CONFIG.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.Set("gorm:tabl_options", "CHARSET=utf8mb4").AutoMigrate(
		&models.Article{}, // 例如 Article 结构体
		&models.Info{},    // 例如 Info 结构体)
	)
	if err != nil {
		os.Exit(0)
	}
}
