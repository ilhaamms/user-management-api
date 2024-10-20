package helper

import (
	"encoding/json"
	"net/http"

	"github.com/ilhaamms/user-management-api/models/response"
)

func ResponseJsonError(w http.ResponseWriter, statusCode int, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response.ErrorResponse{
		StatusCode: statusCode,
		Error:      err,
	})
}

func ResponseJsonSuccess(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response.SuccessResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}

func ResponseJsonSuccessWithPagination(w http.ResponseWriter, statusCode int, message string, page int, limit int, totalPage int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response.ResponseDataWithPagination{
		StatusCode: statusCode,
		Message:    message,
		Pagination: response.Pagination{
			CurrenPage: page,
			TotalPage:  totalPage,
			Limit:      limit,
		},
		Data: data,
	})
}
