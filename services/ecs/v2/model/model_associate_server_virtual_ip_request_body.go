package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type AssociateServerVirtualIpRequestBody struct {
	Nic *AssociateServerVirtualIpOption `json:"nic"`
}

func (o AssociateServerVirtualIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateServerVirtualIpRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateServerVirtualIpRequestBody", string(data)}, " ")
}
