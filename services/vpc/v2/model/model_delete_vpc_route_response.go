package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type DeleteVpcRouteResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpcRouteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcRouteResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpcRouteResponse", string(data)}, " ")
}
