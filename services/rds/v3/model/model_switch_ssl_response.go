package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type SwitchSslResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchSslResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSslResponse struct{}"
	}

	return strings.Join([]string{"SwitchSslResponse", string(data)}, " ")
}
