package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateVolumeTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateVolumeTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateVolumeTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateVolumeTagsResponse", string(data)}, " ")
}
