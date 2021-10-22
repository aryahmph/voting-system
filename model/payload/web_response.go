package payload

type WebResponse struct {
	Code   uint16      `json:"code"`
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}
