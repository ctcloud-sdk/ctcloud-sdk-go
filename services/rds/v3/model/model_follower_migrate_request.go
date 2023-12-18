package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type FollowerMigrateRequest struct {
	NodeId string `json:"nodeId"`

	AzCode string `json:"azCode"`
}

func (o FollowerMigrateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FollowerMigrateRequest struct{}"
	}

	return strings.Join([]string{"FollowerMigrateRequest", string(data)}, " ")
}
