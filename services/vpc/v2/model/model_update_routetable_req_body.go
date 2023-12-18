package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type UpdateRoutetableReqBody struct {
	Routetable *UpdateRouteTableReq `json:"routetable"`
}

func (o UpdateRoutetableReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRoutetableReqBody struct{}"
	}

	return strings.Join([]string{"UpdateRoutetableReqBody", string(data)}, " ")
}
