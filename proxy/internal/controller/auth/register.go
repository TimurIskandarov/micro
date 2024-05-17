package auth

import (
	"encoding/json"
	"net/http"

	"proxy/internal/model"

	"go.uber.org/zap"
)

func (c *controller) Register(w http.ResponseWriter, r *http.Request) {
	req := new(model.RegisterIn)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		c.logger.Error("json decode error", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out := c.service.Register(r.Context(), req)
	
	if out.Status == 409 {
		c.logger.Error("service error", zap.Error(err))
		http.Error(w, out.Message, http.StatusConflict)
		return
	} 
	
	if out.Status == 1 {
		c.logger.Error("json encode error", zap.Error(err))
		http.Error(w, out.Message, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	err = json.NewEncoder(w).Encode(out)
	if err != nil {
		c.logger.Error("json encode error", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
