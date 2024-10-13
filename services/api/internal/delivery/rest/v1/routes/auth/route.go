package auth

import (
	"github.com/kvizyx/glich/services/api/internal/delivery/rest/common"
)

func NewGroup() common.Group {
	return common.Group{
		Prefix: "/auth",
		Routes: []common.Route{
			newAuthSignUp(),
		},
	}
}
