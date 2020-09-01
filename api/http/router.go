package http

import "net/http"

var AppRoutes []RouteSet

type RouteSet struct {
	Prefix    string
	SubRoutes []AppRoute
}

type AppRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Protected   bool
}
