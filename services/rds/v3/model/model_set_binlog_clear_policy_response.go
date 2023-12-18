package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type SetBinlogClearPolicyResponse struct {
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetBinlogClearPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetBinlogClearPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetBinlogClearPolicyResponse", string(data)}, " ")
}
