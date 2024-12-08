package handlers

type ResponseDataDefault struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func NewResponseDataSuccessWithData(data interface{}, message string) ResponseDataDefault {
	return ResponseDataDefault{
		Data:    data,
		Success: true,
		Message: message,
	}
}

func NewResponseDataSuccess(message string) ResponseDataDefault {
	return ResponseDataDefault{
		Success: true,
		Message: message,
	}
}

func NewResponseDataErr(message string) ResponseDataDefault {
	return ResponseDataDefault{
		Success: false,
		Message: message,
	}
}
