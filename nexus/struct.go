package nexus

import "gozealous/service"

type Store struct {
	nativeSvc  service.Native
	railwaySvc service.Railway
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
