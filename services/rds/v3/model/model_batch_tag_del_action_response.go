package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type BatchTagDelActionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchTagDelActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagDelActionResponse struct{}"
	}

	return strings.Join([]string{"BatchTagDelActionResponse", string(data)}, " ")
}
