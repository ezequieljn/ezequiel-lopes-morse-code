package http

import (
	"encoding/json"
	"net/http"

	"github.com/ezequieljn/ezequiel-lopes-morse-code/application"
)

type HttpHandler struct {
	Service application.DecoderService
}

func NewHttpHandler(service application.DecoderService) *HttpHandler {
	return &HttpHandler{Service: service}
}

func (h *HttpHandler) DecodeMorse(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Code string `json:"code"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	decoded, err := h.Service.Decode(request.Code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"decoded": decoded})
}
