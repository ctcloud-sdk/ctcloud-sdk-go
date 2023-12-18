package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type CreateServerGroupRequestBody struct {
	ServerGroup *CreateServerGroupOption `json:"server_group"`
}

func (o CreateServerGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateServerGroupRequestBody", string(data)}, " ")
}
