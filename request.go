package commons

// 分页查询参数
type PagingRequest struct {
	// 一页多少条
	PageSize int64 `json:"pageSize" form:"pageSize"`
	//当前多少页，从1开始
	CurPage int64 `json:"curPage" form:"curPage"`
}
