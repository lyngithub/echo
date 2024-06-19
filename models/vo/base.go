package vo

type FindListParams struct {
	Page      int    `json:"page" doc:"页码从1开始"`
	PageSize  int    `json:"page_size" doc:"每页条数"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func InitPage(vo *FindListParams) {
	if vo.Page <= 0 {
		vo.Page = 1
	}
	if vo.PageSize <= 0 {
		vo.PageSize = 10
	}
}

type Null struct {
}
