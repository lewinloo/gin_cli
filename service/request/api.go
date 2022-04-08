package request

type CreateApiParams struct {
	Path        string `json:"path" `       // api路径
	Description string `json:"description"` // api描述
	ApiGroup    string `json:"apiGroup"`    // api组
	Method      string `json:"method"`      // 请求方法：POST|GET|PUT|PATCH|DELETE
}

type QueryApiListParams struct {
	PageParams
	CreateApiParams
}
