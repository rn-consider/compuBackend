package info

import (
	"github.com/gin-gonic/gin"
	"github.com/rn-consider/compuBackend/models"
	"github.com/rn-consider/compuBackend/service/info/request"
	"net/http"
	"strconv"
)

// CreateInfo 处理创建信息请求的函数
func CreateInfo(c *gin.Context) {
	var req request.CreateInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": "创建信息请求失败，无效的请求参数"})
		return
	}

	info := models.Info{
		Title:   req.Title,
		Content: req.Content,
	}

	if err := models.CreateInfo(&info); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "error": "创建信息请求失败，无法创建信息"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "信息创建成功"})
}

// UpdateInfo 处理更新信息请求的函数
func UpdateInfo(c *gin.Context) {
	var req request.UpdateInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": "更新信息请求失败，无效的请求参数"})
		return
	}

	infoIDStr := c.Param("id") // 假设路由中有 "id" 参数

	infoID, err := strconv.ParseUint(infoIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": "更新信息请求失败，无效的信息 ID"})
		return
	}

	info, err := models.GetInfoByID(uint(infoID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "error": "更新信息请求失败，无法获取信息"})
		return
	}

	// 更新信息的逻辑
	if req.Title != "" {
		info.Title = req.Title
	}
	if req.Content != "" {
		info.Content = req.Content
	}

	if err := models.UpdateInfo(info); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "error": "更新信息请求失败，无法更新信息"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "信息更新成功"})
}

// DeleteInfo 处理删除信息请求的函数
func DeleteInfo(c *gin.Context) {
	infoIDStr := c.Param("id") // 假设路由中有 "id" 参数

	infoID, err := strconv.ParseUint(infoIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": "删除信息请求失败，无效的信息 ID"})
		return
	}

	if err := models.DeleteInfo(uint(infoID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "error": "删除信息请求失败，无法删除信息"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "信息删除成功"})
}

// GetInfo 处理获取单个信息请求的函数
func GetInfo(c *gin.Context) {
	infoIDStr := c.Param("id") // 假设路由中有 "id" 参数

	infoID, err := strconv.ParseUint(infoIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": "获取信息请求失败，无效的信息 ID"})
		return
	}

	info, err := models.GetInfoByID(uint(infoID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "error": "获取信息请求失败，无法获取信息"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "info": info})
}

// GetAllInfo 处理获取所有信息请求的函数
func GetAllInfo(c *gin.Context) {
	infos, err := models.GetAllInfos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "error": "获取所有信息请求失败，无法获取信息列表"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "infos": infos})
}
