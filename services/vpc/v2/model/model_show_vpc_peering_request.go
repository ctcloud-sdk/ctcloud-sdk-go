package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type ShowVpcPeeringRequest struct {
	PeeringId string `json:"peering_id"`
}

func (o ShowVpcPeeringRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcPeeringRequest struct{}"
	}

	return strings.Join([]string{"ShowVpcPeeringRequest", string(data)}, " ")
}
