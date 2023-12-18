package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type DeletePrivateipResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivateipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateipResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateipResponse", string(data)}, " ")
}
