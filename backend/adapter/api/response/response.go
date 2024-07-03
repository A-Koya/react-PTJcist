package response

import (
	"encoding/json"
	"net/http"
)

type responseMessage struct {
	statusCode   int
	ErrorMessage string      `json:"errorMessage"`
	Result       interface{} `json:"result"`
}

func NewErrorMessage(message string, statusCode int) responseMessage {
	return responseMessage{
		statusCode:   statusCode,
		ErrorMessage: message,
		Result:       "",
	}
}

func NewSuccessMessage(result interface{}, statusCode int) responseMessage {
	return responseMessage{
		statusCode:   statusCode,
		ErrorMessage: "",
		Result:       result,
	}
}

func (r responseMessage) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.statusCode)
	return json.NewEncoder(w).Encode(r)
}
