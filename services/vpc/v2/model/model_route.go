package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type Route struct {
	Destination *string `json:"destination,omitempty"`

	Nexthop *string `json:"nexthop,omitempty"`
}

func (o Route) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Route struct{}"
	}

	return strings.Join([]string{"Route", string(data)}, " ")
}
