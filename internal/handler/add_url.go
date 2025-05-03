package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) AddURL(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

}

func validateAddURLPayload(r *http.Request) (AddURLPayload, error) {
	var payload AddURLPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return AddURLPayload{}, err
	}
	return payload, nil
}
