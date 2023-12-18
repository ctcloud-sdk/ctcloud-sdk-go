package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type CreateRouteTableResponse struct {
	Routetable     *RouteTableResp `json:"routetable,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o CreateRouteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRouteTableResponse struct{}"
	}

	return strings.Join([]string{"CreateRouteTableResponse", string(data)}, " ")
}
