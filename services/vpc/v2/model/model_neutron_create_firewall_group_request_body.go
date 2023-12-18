package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateFirewallGroupRequestBody struct {
	FirewallGroup *NeutronCreateFirewallGroupOption `json:"firewall_group"`
}

func (o NeutronCreateFirewallGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallGroupRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallGroupRequestBody", string(data)}, " ")
}
