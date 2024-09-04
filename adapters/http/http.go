package http

import (
	"encoding/json"
	"net/http"

	"github.com/ezequieljn/morse-code/application"
)

type HttpHandler struct {
	Service application.DecoderService
}

func DecodeHandler(version, space string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			Code string `json:"code"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			responseError(w, err, http.StatusBadRequest)
			return
		}
		morse, err := application.NewDecoderFactory(version, space)
		if err != nil {
			responseError(w, err, http.StatusUnprocessableEntity)
			return
		}
		value, err := morse.Decode(request.Code)
		if err != nil {
			responseError(w, err, http.StatusUnprocessableEntity)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"decoded": value})
	}

}
func responseError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
