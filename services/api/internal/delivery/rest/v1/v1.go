package v1

import (
	"github.com/kvizyx/glich/services/api/internal/delivery/rest/common"
	"github.com/kvizyx/glich/services/api/internal/delivery/rest/v1/routes/auth"
)

// @title Glich REST API
// @version 1.0
// @description Current v1 REST API for Glich
// @termsOfService TODO
//
// @license.name Apache 2.0
// @license.url https://github.com/glichapp/REST/blob/dev/LICENSE.md
//
// @host localhost:8080
// @BasePath /v1
// @schemes http https
// @query.collection.format multi
func NewGroup(router common.Router) common.Group {
	return common.Group{
		Prefix: "/v1",
		Children: []common.Group{
			auth.NewGroup(),
		},
	}
}
