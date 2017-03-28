package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"io"
	"secureStore/models"
)

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

func DeleteSubmissionAction(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	err := models.DeleteSubmission(vars["submissionId"])

	if err != nil {
		buildResponse(w, APIPayload{"status": "error", "message": err.Error()}, http.StatusBadRequest)
		return
	}

	buildResponse(w, APIPayload{"status": "success"}, http.StatusOK)
}
