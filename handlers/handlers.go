package handlers

import (
	"net/http"
	"encoding/json"
	"errors"
	"fmt"
)

type APIPayload map[string]interface{}

func buildResponse(w http.ResponseWriter, payload APIPayload, statusCode int) {
	j, _ := json.Marshal(payload)

	w.Header().Set("Content-Type",	"application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	w.Write(j)
}

func checkRequiredFields(r *http.Request, keys ...string) (error) {
	r.ParseForm()

	for _, key := range keys {
		if len(r.Form[key]) == 0 {
			return errors.New(fmt.Sprintf("%s is required.", key))
		}
	}

	return nil
}

func IndexAction(w http.ResponseWriter, request *http.Request) {
	payload := APIPayload{
		"status": "success",
		"data": "Index",
	}

	buildResponse(w, payload, http.StatusOK)
}