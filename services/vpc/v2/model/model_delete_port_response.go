package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type DeletePortResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePortResponse struct{}"
	}

	return strings.Join([]string{"DeletePortResponse", string(data)}, " ")
}
