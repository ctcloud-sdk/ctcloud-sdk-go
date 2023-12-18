package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type CreateOrUpdateTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateOrUpdateTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateTagsResponse", string(data)}, " ")
}
