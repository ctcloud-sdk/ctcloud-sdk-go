package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type AttachEipResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AttachEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachEipResponse struct{}"
	}

	return strings.Join([]string{"AttachEipResponse", string(data)}, " ")
}
