package models

type OperationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type OperationResponseSucess[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}

func GetOperationErrorResponse(message string) OperationResponse {
	return OperationResponse{
		Status:  "Error",
		Message: message,
	}
}

func GetOperationFailureResponse(message string) OperationResponse {
	return OperationResponse{
		Status:  "Fail",
		Message: message,
	}
}

func GetOperationSuccessResponse[T any](data T) OperationResponseSucess[T] {
	return OperationResponseSucess[T]{
		Status: "Successful",
		Data:   data,
	}
}
