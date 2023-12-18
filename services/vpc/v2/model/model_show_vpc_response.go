package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type ShowVpcResponse struct {
	Vpc            *Vpc `json:"vpc,omitempty"`
	HttpStatusCode int  `json:"-"`
}

func (o ShowVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcResponse struct{}"
	}

	return strings.Join([]string{"ShowVpcResponse", string(data)}, " ")
}
