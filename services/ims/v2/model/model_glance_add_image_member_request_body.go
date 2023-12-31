package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type GlanceAddImageMemberRequestBody struct {
	Member string `json:"member"`
}

func (o GlanceAddImageMemberRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceAddImageMemberRequestBody struct{}"
	}

	return strings.Join([]string{"GlanceAddImageMemberRequestBody", string(data)}, " ")
}
