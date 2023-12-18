package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type GlanceShowImageSchemasRequest struct {
}

func (o GlanceShowImageSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageSchemasRequest struct{}"
	}

	return strings.Join([]string{"GlanceShowImageSchemasRequest", string(data)}, " ")
}
