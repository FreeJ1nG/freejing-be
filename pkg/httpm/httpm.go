package httpm

import (
	"encoding/json"
	"net/http"
)

type Response[T interface{}] struct {
	Data       T      `json:"data,omitempty"`
	StatusCode int    `json:"statusCode"`
	Success    bool   `json:"success"`
	Error      string `json:"errors,omitempty"`
}

func MakeErrorResponse(w http.ResponseWriter, httpStatus int, err error) []byte {
	w.WriteHeader(httpStatus)
	response := Response[interface{}]{Data: nil, StatusCode: httpStatus, Success: false, Error: err.Error()}

	json, _ := json.Marshal(response)

	return json
}

func MakeSuccessResponse[T interface{}](w http.ResponseWriter, httpStatus int, data interface{}) []byte {
	var response interface{}
	if data == nil {
		w.WriteHeader(httpStatus)
		response = Response[T]{StatusCode: httpStatus, Success: true}
	} else {
		switch v := data.(type) {
		case T:
			w.WriteHeader(httpStatus)
			response = Response[T]{Data: v, StatusCode: httpStatus, Success: true}
		case []T:
			w.WriteHeader(httpStatus)
			if len(v) == 0 {
				response = Response[[]T]{Data: []T{}, StatusCode: httpStatus, Success: true}
			} else {
				response = Response[[]T]{Data: v, StatusCode: httpStatus, Success: true}
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
			response = Response[T]{StatusCode: http.StatusInternalServerError, Success: false, Error: "invalid data type"}
		}
	}

	json, _ := json.Marshal(response)

	return json
}
