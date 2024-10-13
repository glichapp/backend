package common

import (
	"github.com/gin-gonic/gin"
)

type APIError struct {
	StatusCode int
	Error      string `json:"error"`
}

type Router = gin.IRouter

type (
	Handler      func(ctx *gin.Context) *APIError
	ErrorHandler func(ctx *gin.Context, err *APIError)
)

type Group struct {
	Prefix   string
	Routes   []Route
	Children []Group
}

type Route struct {
	Options RouteOptions
	Handler Handler
}

type RouteOptions struct {
	Method       string
	RelativePath string
	Middlewares  []Handler
}

func DefaultErrorHandler(ctx *gin.Context, err *APIError) {
	if err == nil {
		return
	}

	ctx.JSON(err.StatusCode, err)
}

func (g Group) Register(eh ErrorHandler, root Router) {
	router := root.Group(g.Prefix)

	for _, route := range g.Routes {
		handlers := make([]Handler, len(route.Options.Middlewares)+1)
		copy(handlers[:len(handlers)-1], route.Options.Middlewares)

		handlers[len(handlers)-1] = route.Handler

		router.Handle(
			route.Options.Method,
			route.Options.RelativePath,
			toGinHandlers(eh, handlers)...,
		)
	}

	for _, child := range g.Children {
		group := Group{
			Prefix:   child.Prefix,
			Routes:   child.Routes,
			Children: child.Children,
		}

		group.Register(eh, router)
	}
}

func toGinHandlers(eh ErrorHandler, handlers []Handler) []gin.HandlerFunc {
	ginHandlers := make([]gin.HandlerFunc, len(handlers))

	for index, handler := range handlers {
		ginHandlers[index] = func(ctx *gin.Context) {
			if err := handler(ctx); err != nil {
				eh(ctx, err)
			}
		}
	}

	return ginHandlers
}
