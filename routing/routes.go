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
	Route{
		"GETSubmission",
		"/submission/{submissionId}",
		"GET",
		handlers.GetSubmissionAction,
	},
	Route{
		"POSTSubmission",
		"/submission/",
		"POST",
		handlers.PostSubmissionAction,
	},
	Route{
		"GETSubmissions",
		"/submissions/",
		"GET",
		handlers.GetSubmissionsAction,
	},
}
