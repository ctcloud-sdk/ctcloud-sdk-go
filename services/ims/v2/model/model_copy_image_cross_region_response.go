package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type CopyImageCrossRegionResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyImageCrossRegionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyImageCrossRegionResponse struct{}"
	}

	return strings.Join([]string{"CopyImageCrossRegionResponse", string(data)}, " ")
}
