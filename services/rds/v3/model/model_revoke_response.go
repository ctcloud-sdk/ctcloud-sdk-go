package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type RevokeResponse struct {
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RevokeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokeResponse struct{}"
	}

	return strings.Join([]string{"RevokeResponse", string(data)}, " ")
}
