package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type DeleteVpcTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpcTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpcTagResponse", string(data)}, " ")
}
