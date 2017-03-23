package handlers

import (
	"net/http"
	"secureStore/models"
)

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
