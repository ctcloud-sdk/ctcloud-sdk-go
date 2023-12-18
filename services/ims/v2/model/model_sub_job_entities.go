package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type SubJobEntities struct {
	ImageId *string `json:"image_id,omitempty"`

	ImageName *string `json:"image_name,omitempty"`
}

func (o SubJobEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJobEntities struct{}"
	}

	return strings.Join([]string{"SubJobEntities", string(data)}, " ")
}
