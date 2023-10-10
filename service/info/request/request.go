package request

// CreateInfoRequest 结构体用于创建信息的请求
type CreateInfoRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// UpdateInfoRequest 结构体用于更新信息的请求
type UpdateInfoRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// DeleteInfoRequest 结构体用于删除信息的请求
type DeleteInfoRequest struct {
	ID uint `json:"id" binding:"required"`
}
