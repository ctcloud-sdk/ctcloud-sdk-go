package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type SubnetList struct {
	Id string `json:"id"`
}

func (o SubnetList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubnetList struct{}"
	}

	return strings.Join([]string{"SubnetList", string(data)}, " ")
}
