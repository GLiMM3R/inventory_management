package types

type Response struct {
	Status   int         `json:"status"`
	Messages string      `json:"messages"`
	Data     interface{} `json:"data"`
	Total    *int64      `json:"total"`
}
