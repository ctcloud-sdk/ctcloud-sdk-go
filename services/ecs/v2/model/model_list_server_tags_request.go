package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type ListServerTagsRequest struct {
}

func (o ListServerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerTagsRequest struct{}"
	}

	return strings.Join([]string{"ListServerTagsRequest", string(data)}, " ")
}
