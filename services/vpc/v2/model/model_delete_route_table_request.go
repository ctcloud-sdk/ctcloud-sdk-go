package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type DeleteRouteTableRequest struct {
	RoutetableId string `json:"routetable_id"`
}

func (o DeleteRouteTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRouteTableRequest struct{}"
	}

	return strings.Join([]string{"DeleteRouteTableRequest", string(data)}, " ")
}
