package routes

import "net/http"

type ApiRestRouteSet struct {
	Prefix    string
	SubRoutes []ApiRestRoute
}

type ApiRestRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Protected   bool
}
