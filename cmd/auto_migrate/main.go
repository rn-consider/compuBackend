package main

import (
	"database/sql"
	"github.com/rn-consider/compuBackend/cmd/initialize"
	"github.com/rn-consider/compuBackend/config"
	"github.com/rn-consider/compuBackend/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	config.GVA_VP = initialize.Viper()
	config.GVA_DB = initialize.Gorm() // gorm连接数据库

	if config.GVA_DB != nil {
		// 使用 AutoMigrate 方法创建表
		config.GVA_DB.AutoMigrate(&models.Info{})

		// 清空现有的数据表以备份数据
		config.GVA_DB.Exec("TRUNCATE TABLE infos;")

		// 插入多条初始数据到 infos 表
		initData := []models.Info{
			{
				ID:      1,
				Title:   "学会简介",
				Content: "这里是学会简介",
			},
			{
				ID:      2,
				Title:   "学会简史",
				Content: "这里是学会简史",
			},
			{
				ID:      3,
				Title:   "学会荣誉",
				Content: "这里是学会荣誉",
			},
			{
				ID:      4,
				Title:   "章程制度",
				Content: "这里是章程制度",
			},
			{
				ID:      5,
				Title:   "组织机构",
				Content: "这里是组织机构",
			},
		}

		for _, info := range initData {
			// 检查是否已经存在相应ID的记录
			var existingInfo models.Info
			result := config.GVA_DB.First(&existingInfo, info.ID)
			if result.Error == gorm.ErrRecordNotFound {
				// 如果记录不存在，插入新记录
				config.GVA_DB.Create(&info)
			}
		}

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
