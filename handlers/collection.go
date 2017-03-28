package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"secureStore/models"
)

func GetCollectionAction(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	collection, err := models.GetCollection(vars["collectionId"])

	if err != nil {
		buildResponse(w, APIPayload{"status": "error", "message": "No collections found."}, http.StatusNotFound)
		return
	}

	buildResponse(
		w,
		APIPayload{
			"status": "success",
			"data": collection,
		},
		http.StatusOK,
	)
}

func PostCollectionAction(w http.ResponseWriter, request *http.Request) {
	err := checkRequiredFields(request, "title", "ownerId", "recipients")

	if err != nil {
		buildResponse(w, APIPayload{"status": "error", "message": err.Error()}, http.StatusBadRequest)
		return
	}

	title      := request.PostFormValue("title")
	ownerId    := request.PostFormValue("ownerId")
	recipients := request.PostFormValue("recipients")

	data, err := models.InsertCollection(title, ownerId, recipients)

	if err != nil {
		buildResponse(w, APIPayload{"status": "error", "message": err.Error()}, http.StatusBadRequest)
		return
	}

	buildResponse(w, APIPayload{"status": "success", "data": data}, http.StatusOK)
}

func DeleteCollectionAction(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	err := models.DeleteCollection(vars["collectionId"])

	if err != nil {
		buildResponse(w, APIPayload{"status": "error", "message": err.Error()}, http.StatusBadRequest)
		return
	}

	buildResponse(w, APIPayload{"status": "success"}, http.StatusOK)
}


