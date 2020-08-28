package base

/**
响应列表模板
*/
type ResponseListTem struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

/**
响应列表含头列表
*/
type ResponseHeadListTem struct {
	Total int         `json:"total"`
	Head  interface{} `json:"head"`
	List  interface{} `json:"list"`
}
