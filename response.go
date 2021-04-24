package commons

import "math"

// 通用的返回模型
type ActionResponse struct {
	// 是否成功
	Success bool `json:"success"`
	// 错误码
	ErrCode int `json:"errCode,omitempty"`
	// 错误内容
	ErrMessage string `json:"errMessage,omitempty"`
	// 返回值
	Data interface{} `json:"data,omitempty"`
}

// 分页返回模型
type PagingResponse struct {
	PagingRequest
	// 总页数
	TotalPage int64 `json:"totalPage"`
	// 总记录数
	TotalCount int64 `json:"totalCount"`
	// 分页实体数据
	Data []interface{} `json:"data,omitempty"`
}

// 计算总页数
func (resp *PagingResponse) Compute() {
	if resp.TotalCount == 0 {
		return
	}
	resp.TotalPage = int64(math.Ceil(float64(resp.TotalCount) / float64(resp.PageSize)))
}
