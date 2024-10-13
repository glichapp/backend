package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kvizyx/glich/services/api/internal/delivery/rest/common"
)

type authSignUp struct {
	// Depenencies
}

func newAuthSignUp() common.Route {
	route := authSignUp{
		// Dependencies
	}

	return common.Route{
		Options: common.RouteOptions{
			Method:       http.MethodPost,
			RelativePath: "/sign-up",
		},
		Handler: route.Handler,
	}
}

func (r authSignUp) Handler(ctx *gin.Context) *common.APIError {
	return nil
}
