package dto

type ResponseDto struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseSuccess(data interface{}) ResponseDto {
	return ResponseDto{
		Success: true,
		Message: "operation successfully",
		Error:   "",
		Data:    data,
	}
}

func ResponseFailed(err error) ResponseDto {
	return ResponseDto{
		Success: false,
		Message: "operation error",
		Error:   err.Error(),
		Data:    nil,
	}
}
