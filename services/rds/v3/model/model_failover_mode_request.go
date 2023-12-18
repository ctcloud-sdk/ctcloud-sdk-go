package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type FailoverModeRequest struct {
	Mode string `json:"mode"`
}

func (o FailoverModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailoverModeRequest struct{}"
	}

	return strings.Join([]string{"FailoverModeRequest", string(data)}, " ")
}
