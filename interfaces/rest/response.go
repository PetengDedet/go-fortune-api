package rest

import "net/http"

type ErrorPayload struct {
	Message          string `json:"message"`
	Reason           string `json:"reason"`
	ErrorUserMessage string `json:"error_user_msg"`
	ErrorUserTitle   string `json:"error_user_title"`
}

type ErrorResponse struct {
	Status  int64         `json:"status"`
	Message string        `json:"message"`
	Data    interface{}   `json:"data"`
	Error   *ErrorPayload `json:"error"`
}

func NotFoundResponse(data interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status:  http.StatusNotFound,
		Message: "failed",
		Error: &ErrorPayload{
			Message:          "Page doesn't exists",
			Reason:           "not_found",
			ErrorUserTitle:   "",
			ErrorUserMessage: "",
		},
	}
}

func InternalErrorResponse(data interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: "failed",
		Error: &ErrorPayload{
			Message:          "internal server error",
			Reason:           "Something went wrong",
			ErrorUserTitle:   "",
			ErrorUserMessage: "",
		},
	}
}

type Response struct {
	Status  int64       `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(data interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    data,
	}
}
