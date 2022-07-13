package input

//分页
type PaginationInput struct {
	PageIndex int `form:"page_index" json:"page_index" comment:"页码" binding:"required"`
	PageSize  int `form:"page_size" json:"page_size" comment:"页容量" binding:"required"`
}

//可选分页
type OptionalPaginationInput struct {
	PageIndex int `form:"page_index" json:"page_index" comment:"页码"`
	PageSize  int `form:"page_size" json:"page_size" comment:"页容量"`
}

//日期范围
type DateInput struct {
	StartDate int64 `form:"start_date" json:"start_date" comment:"开始时间" binding:"required"`
	EndDate   int64 `form:"end_date" json:"end_date" comment:"结束时间" binding:"required"`
}

//可选日期范围
type OptionalDateInput struct {
	StartDate int64 `form:"start_date" json:"start_date" comment:"开始时间"`
	EndDate   int64 `form:"end_date" json:"end_date" comment:"结束时间"`
}

//可选查询参数
type OptionalQuery struct {
	Query string `form:"query" json:"query" comment:"查询参数"`
}
