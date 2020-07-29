package response

type MyResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}
