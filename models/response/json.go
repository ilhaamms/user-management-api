package response

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

type SuccessResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type ResponseDataWithPagination struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Pagination
	Data any `json:"data"`
}
