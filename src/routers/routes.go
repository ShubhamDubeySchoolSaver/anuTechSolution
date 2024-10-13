package router

import (
	"Skool_Saver/src/handlers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{
		"404 Not Found",
		"GET, POST",
		"/",
		ThrowBlankResponse,
	},
	Route{
		"InsertUserDetails",
		"POST",
		"/insertUserDetails",
		handlers.InsertUserDetailsHandler,
	},
	Route{
		"UpsertCandidateDetails",
		"POST",
		"/upsertCandidateDetails",
		handlers.UpsertCandidateDetails,
	},
	Route{
		"InsertQuotation",
		"POST",
		"/insertQuotation",
		handlers.InsertQuotationHandler,
	},
	Route{
		"Login",
		"POST",
		"/login",
		handlers.LoginHandler,
	},
	Route{
		"FetchGetInTouch",
		"GET",
		"/fetchGetInTouchDetails",
		handlers.FetchGetInTouchHandlers,
	},
	Route{
		"FetchJobAppliedCandidates",
		"GET",
		"/fetchJobAppliedCandidatesDetails",
		handlers.FetchJobAppliedCandidatesHandlers,
	},
	Route{
		"FetchConsultations",
		"GET",
		"/fetchConsultations",
		handlers.FetchConsultationsHandlers,
	},
	Route{
		"UploadJobs",
		"POST",
		"/uploadJob",
		handlers.UploadJobHandlers,
	},
	Route{
		"FetchJobsDetails",
		"GET",
		"/fetchJobsDetails",
		handlers.FetchJobsHandlers,
	},
}
