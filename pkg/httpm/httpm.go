package httpm

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Success    bool        `json:"success"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
}

func MakeErrorResponse(w http.ResponseWriter, httpStatus int, err error) []byte {
	w.WriteHeader(httpStatus)
	response := Response{Data: nil, StatusCode: httpStatus, Success: false, Error: err.Error()}

	json, _ := json.Marshal(response)

	return json
}

func MakeSuccessResponse[T interface{}](w http.ResponseWriter, httpStatus int, data interface{}) []byte {
	var response interface{}
	if data == nil {
		w.WriteHeader(httpStatus)
		response = Response{StatusCode: httpStatus, Success: true}
	} else {
		switch v := data.(type) {
		case T:
			w.WriteHeader(httpStatus)
			response = Response{Data: v, StatusCode: httpStatus, Success: true}
		case []T:
			w.WriteHeader(httpStatus)
			if len(v) == 0 {
				response = Response{Data: []T{}, StatusCode: httpStatus, Success: true}
			} else {
				response = Response{Data: v, StatusCode: httpStatus, Success: true}
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
			response = Response{StatusCode: http.StatusInternalServerError, Success: false, Error: "invalid data type"}
		}
	}

	json, _ := json.Marshal(response)

	return json
}
