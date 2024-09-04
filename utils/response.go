package utils

type Response struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
	Error any `json:"error,omitempty"`
}

func BuildResponseSuccess(message string, data any) Response {
	res := Response{
		Status: true,
		Message: message,
		Data: data,
	}

	return res
}

func BuildResponseFailed(message string, err string, data any) Response {
	res := Response{
		Status: false,
		Message: message,
		Error: err,
		Data: data,
	}

	return res
}