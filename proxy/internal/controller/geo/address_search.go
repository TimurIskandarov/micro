package geo

import (
	"encoding/json"
	"net/http"

	"proxy/internal/model"

	"go.uber.org/zap"
)

func (c *controller) AddressSearch(w http.ResponseWriter, r *http.Request) {
	var req model.SearchIn

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		c.logger.Error("json decode error", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := c.service.AddressSearch(r.Context(), req)
	if err != nil {
		c.logger.Error("service error", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	err = json.NewEncoder(w).Encode(out.Addresses)
	if err != nil {
		c.logger.Error("json encode error", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
