package models

import (
	"github.com/rn-consider/compuBackend/config"
	"time"
)

// Info 结构体映射到 "infos" 表
type Info struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Content   string    `gorm:"type:longtext;not null" json:"content"`
	CreatedAt time.Time `gorm:"type:timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null" json:"updated_at"`
}

// TableName 指定表名为 "infos"
func (Info) TableName() string {
	return "infos"
}

// CreateInfo 创建信息
func CreateInfo(info *Info) error {
	if err := config.GVA_DB.Create(info).Error; err != nil {
		return err
	}
	return nil
}

// DeleteInfo 根据ID删除信息
func DeleteInfo(id uint) error {
	if err := config.GVA_DB.Where("id = ?", id).Delete(&Info{}).Error; err != nil {
		return err
	}
	return nil
}

// UpdateInfo 更新信息
func UpdateInfo(info *Info) error {
	if err := config.GVA_DB.Save(info).Error; err != nil {
		return err
	}
	return nil
}

// GetAllInfos 获取所有信息
func GetAllInfos() ([]*Info, error) {
	var infos []*Info
	if err := config.GVA_DB.Find(&infos).Error; err != nil {
		return nil, err
	}
	return infos, nil
}

// GetInfoByID 根据ID获取信息
func GetInfoByID(id uint) (*Info, error) {
	var info Info
	if err := config.GVA_DB.Where("id = ?", id).First(&info).Error; err != nil {
		return nil, err
	}
	return &info, nil
}
