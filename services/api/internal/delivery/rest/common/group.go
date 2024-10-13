package common

import (
	"github.com/gin-gonic/gin"
)

type Router = gin.IRouter

type Handler func(ctx *gin.Context) *APIError

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

func (r Route) Register(eh ErrorHandler, router Router) {
	handlers := make([]Handler, len(r.Options.Middlewares)+1)
	copy(handlers[:len(handlers)-1], r.Options.Middlewares)

	handlers[len(handlers)-1] = r.Handler

	router.Handle(
		r.Options.Method,
		r.Options.RelativePath,
		toGinHandlers(eh, handlers)...,
	)
}

func (g Group) Register(eh ErrorHandler, root Router) {
	router := root.Group(g.Prefix)

	for _, route := range g.Routes {
		route.Register(eh, router)
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
