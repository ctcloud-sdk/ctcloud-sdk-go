package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteVolumeTagsRequest struct {
	VolumeId string `json:"volume_id"`

	Body *BatchDeleteVolumeTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteVolumeTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteVolumeTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteVolumeTagsRequest", string(data)}, " ")
}
