package organization_info

import (
	"github.com/gin-gonic/gin"
	r "github.com/rn-consider/compuBackend/service/organization_info/request"
	"net/http"
)

// CreateOrganization 处理创建组织请求的函数
func CreateOrganization(c *gin.Context) {
	var request r.CreateOrganizationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Organization created successfully"})
}

// UpdateOrganization 处理更新组织请求的函数
func UpdateOrganization(c *gin.Context) {
	var request UpdateOrganizationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 处理更新组织的逻辑
	// ...
	c.JSON(http.StatusOK, gin.H{"message": "Organization updated successfully"})
}

// deleteOrganization 处理删除组织请求的函数
func DeleteOrganization(c *gin.Context) {
	var request DeleteOrganizationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 处理删除组织的逻辑
	// ...
	c.JSON(http.StatusOK, gin.H{"message": "Organization deleted successfully"})
}

// GetOrganization  处理获取组织请求的函数
func GetOrganization(c *gin.Context) {
	var request GetOrganizationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 处理获取组织的逻辑
	// ...
	c.JSON(http.StatusOK, gin.H{"message": "Organization retrieved successfully"})
}
