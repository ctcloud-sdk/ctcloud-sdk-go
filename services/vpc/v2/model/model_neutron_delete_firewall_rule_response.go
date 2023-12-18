package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type NeutronDeleteFirewallRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteFirewallRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteFirewallRuleResponse struct{}"
	}

	return strings.Join([]string{"NeutronDeleteFirewallRuleResponse", string(data)}, " ")
}
