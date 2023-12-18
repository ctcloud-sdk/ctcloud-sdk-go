package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type ShowSecurityGroupResponse struct {
	SecurityGroup  *SecurityGroup `json:"security_group,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityGroupResponse", string(data)}, " ")
}
