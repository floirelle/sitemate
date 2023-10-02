package response

type BaseResponse struct {
	StatusCode int `json:"statusCode"`
	Data       any `json:"data"`
}

func Json(statusCode int, data any) BaseResponse {
	return BaseResponse{
		StatusCode: statusCode,
		Data:       data,
	}
}
