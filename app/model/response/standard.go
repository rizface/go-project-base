package response

type Standard struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   interface{} `json:"data"`
}
