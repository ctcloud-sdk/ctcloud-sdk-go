package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type GlanceListImageMemberSchemasRequest struct {
}

func (o GlanceListImageMemberSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImageMemberSchemasRequest struct{}"
	}

	return strings.Join([]string{"GlanceListImageMemberSchemasRequest", string(data)}, " ")
}
