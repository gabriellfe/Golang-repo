package routes

import (
	"gabriellfe/handler"
	"gabriellfe/middleware"
	"net/http"
)

type RouteDTO struct {
	Route      string
	Handler    func(http.ResponseWriter, *http.Request)
	Middleware []func(next http.HandlerFunc) http.HandlerFunc
}

var Routes = []RouteDTO{
	RouteDTO{Route: "GET actuator/live", Handler: handler.LiveHandler, Middleware: []func(next http.HandlerFunc) http.HandlerFunc{middleware.ApplicationJSON, middleware.SchemaValidatorMiddleware, middleware.LoggingMiddleware}},
	RouteDTO{Route: "GET /actuator/refresh", Handler: handler.RefreshHandler, Middleware: []func(next http.HandlerFunc) http.HandlerFunc{middleware.ApplicationJSON, middleware.LoggingMiddleware}},
}

func Setup(router *http.ServeMux) {
	for _, item := range Routes {
		router.HandleFunc(item.Route, callRoute(item.Handler, item.Middleware))
	}

	for _, item := range UserRoutes {
		router.HandleFunc(item.Route, callRoute(item.Handler, item.Middleware))
	}
}

func callRoute(next http.HandlerFunc, middlewares []func(next http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	// Execute Middlewares for route
	for _, item := range middlewares {
		next = item(next)
	}
	// Execute final Handler
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
}
