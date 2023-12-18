package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type ListVolumeTagsRequest struct {
}

func (o ListVolumeTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumeTagsRequest struct{}"
	}

	return strings.Join([]string{"ListVolumeTagsRequest", string(data)}, " ")
}
