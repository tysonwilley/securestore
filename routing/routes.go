package routing

import (
	"secureStore/handlers"
	"net/http"
)

type Route struct{
	Name, Path, Method string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"/",
		"GET",
		handlers.IndexAction,
	},

	// Submissions
	// ======================================
	Route{
		"GETSubmission",
		"/submission/{submissionId}",
		"GET",
		handlers.GetSubmissionAction,
	},
	Route{
		"POSTSubmission",
		"/submission/{collectionId}",
		"POST",
		handlers.PostSubmissionAction,
	},
	Route{
		"DELETESubmission",
		"/submission/{submissionId}",
		"DELETE",
		handlers.DeleteSubmissionAction,
	},
	Route{
		"GETSubmissions",
		"/submissions/",
		"GET",
		handlers.GetSubmissionsAction,
	},

	// Collections
	// ======================================
	Route{
		"GETCollection",
		"/collection/{collectionId}",
		"GET",
		handlers.GetCollectionAction,
	},
	Route{
		"POSTCollection",
		"/collection/",
		"POST",
		handlers.PostCollectionAction,
	},
	Route{
		"DELETECollection",
		"/collection/{collectionId}",
		"DELETE",
		handlers.DeleteCollectionAction,
	},
}