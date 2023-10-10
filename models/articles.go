package models

import (
	"github.com/rn-consider/compuBackend/dao"
	"time"
)

// Article 结构体映射到 "articles" 表
type Article struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	PublishTime time.Time `gorm:"type:timestamp;not null" json:"publish_time"`
	Author      string    `gorm:"type:varchar(255);not null" json:"author"`
	Content     string    `gorm:"type:text;not null" json:"content"`
}

// TableName 指定表名为 "articles"
func (Article) TableName() string {
	return "articles"
}

/*
   Article的增删改查
*/

// CreateArticle 创建文章
func CreateArticle(article *Article) error {
	if err := dao.DB.Create(article).Error; err != nil {
		return err
	}
	return nil
}

// DeleteArticle 根据主键删除文章
func DeleteArticle(id uint) error {
	if err := dao.DB.Where("id = ?", id).Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}

// UpdateArticle 更新文章
func UpdateArticle(article *Article) error {
	if err := dao.DB.Save(article).Error; err != nil {
		return err
	}
	return nil
}

// GetAllArticles 获取所有文章
func GetAllArticles() ([]*Article, error) {
	var articles []*Article
	if err := dao.DB.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

// GetArticleByID 根据ID获取文章
func GetArticleByID(id uint) (*Article, error) {
	var article Article
	if err := dao.DB.Where("id = ?", id).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}
