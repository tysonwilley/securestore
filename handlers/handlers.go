package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"secureStore/models"
	"encoding/json"
	"io/ioutil"
	"io"
)

type APIPayload map[string]interface{}

func buildResponse(w http.ResponseWriter, payload APIPayload, statusCode int) {
	j, _ := json.Marshal(payload)

	w.Header().Set("Content-Type",	"application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	w.Write(j)
}

func IndexAction(w http.ResponseWriter, request *http.Request) {
	payload := APIPayload{
		"status": "success",
		"data": "Index",
	}

	buildResponse(w, payload, http.StatusOK)
}

func GetSubmissionAction(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	submission, err := models.GetSubmission(vars["submissionId"])

	if err != nil {
		buildResponse(w, APIPayload{"status": "error", "message": "No submissions found."}, http.StatusNotFound)
		return
	}

	buildResponse(
		w,
		APIPayload{
			"status": "success",
			"data": submission,
		},
		http.StatusOK,
	)
}

func PostSubmissionAction(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	collectionId := vars["collectionId"]

	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
	if err != nil {
		buildResponse(w, APIPayload{"status": "error", "message": err.Error()}, http.StatusBadRequest)
		return
	}

	data, err := models.InsertSubmission(body, collectionId)

	if err != nil {
		buildResponse(w, APIPayload{"status": "error", "message": err.Error()}, http.StatusBadRequest)
		return
	}

	buildResponse(w, APIPayload{"status": "success", "data": data}, http.StatusOK)
}

func GetSubmissionsAction(w http.ResponseWriter, request *http.Request) {
	collectionId := request.FormValue("collectionId")

	if len(collectionId) == 0 {
		buildResponse(w, APIPayload{"status": "error", "message": "collectionId is required."}, http.StatusBadRequest)
		return
	}

	submissions, err := models.GetSubmissions(collectionId)

	if err != nil {
		buildResponse(w, APIPayload{"status": "error", "message": err.Error()}, http.StatusInternalServerError)
		return
	}

	buildResponse(w, APIPayload{"status": "success", "data": submissions}, http.StatusOK)
}