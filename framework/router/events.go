package router

import (
	"net/http"

	"flamingo.me/flamingo/framework/web"
)

type (
	// OnRequestEvent contains the bare request
	OnRequestEvent struct {
		ResponseWriter http.ResponseWriter
		Request        *http.Request
	}

	// OnResponseEvent is the event associated to OnResponse
	OnResponseEvent struct {
		Response       web.Response
		Request        *http.Request
		ResponseWriter http.ResponseWriter
	}

	// OnFinishEvent is the event object associated to OnFinish
	OnFinishEvent struct {
		ResponseWriter http.ResponseWriter
		Request        *http.Request
		Error          error
	}
)
