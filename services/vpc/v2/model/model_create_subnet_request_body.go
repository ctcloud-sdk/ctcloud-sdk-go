package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type CreateSubnetRequestBody struct {
	Subnet *CreateSubnetOption `json:"subnet"`
}

func (o CreateSubnetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubnetRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSubnetRequestBody", string(data)}, " ")
}
