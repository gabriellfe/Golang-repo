package routes

import (
	"gabriellfe/handler"
	"gabriellfe/middleware"
	"net/http"
)

var UserRoutes = []RouteDTO{
	RouteDTO{Route: "GET /user/{id}", Handler: handler.GetUserHandler, Middleware: []func(next http.HandlerFunc) http.HandlerFunc{middleware.SchemaValidatorMiddleware, middleware.ApplicationJSON, middleware.LoggingMiddleware}},
	RouteDTO{Route: "POST /user", Handler: handler.PostUserHandler, Middleware: []func(next http.HandlerFunc) http.HandlerFunc{middleware.SchemaValidatorMiddleware, middleware.ApplicationJSON, middleware.LoggingMiddleware}},
}
