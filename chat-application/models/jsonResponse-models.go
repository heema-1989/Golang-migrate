package models

type JsonResponse struct {
	Status  int
	Message string
	Data    interface{}
}

func Response(status int, message string, data interface{}, resp JsonResponse) JsonResponse {
	resp.Status = status
	resp.Message = message
	resp.Data = data
	return resp
}
