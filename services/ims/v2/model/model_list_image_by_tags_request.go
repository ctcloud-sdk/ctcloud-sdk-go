package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type ListImageByTagsRequest struct {
	Body *ListImageByTagsRequestBody `json:"body,omitempty"`
}

func (o ListImageByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListImageByTagsRequest", string(data)}, " ")
}
