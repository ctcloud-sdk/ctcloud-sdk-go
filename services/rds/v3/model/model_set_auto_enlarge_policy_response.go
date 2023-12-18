package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type SetAutoEnlargePolicyResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetAutoEnlargePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoEnlargePolicyResponse struct{}"
	}

	return strings.Join([]string{"SetAutoEnlargePolicyResponse", string(data)}, " ")
}
