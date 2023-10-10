package request

// CreateOrganizationRequest 结构体用于创建组织的请求
type CreateOrganizationRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// UpdateOrganizationRequest 结构体用于更新组织的请求
type UpdateOrganizationRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DeleteOrganizationRequest 结构体用于删除组织的请求
type DeleteOrganizationRequest struct {
	ID uint `json:"id" binding:"required"`
}
