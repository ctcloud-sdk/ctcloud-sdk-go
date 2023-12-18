package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type CreateDataImageRequest struct {
	Body *CreateDataImageRequestBody `json:"body,omitempty"`
}

func (o CreateDataImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataImageRequest struct{}"
	}

	return strings.Join([]string{"CreateDataImageRequest", string(data)}, " ")
}
