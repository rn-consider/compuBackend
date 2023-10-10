package response

// OrganizationResponse 结构体用于表示组织信息的响应
type OrganizationResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
