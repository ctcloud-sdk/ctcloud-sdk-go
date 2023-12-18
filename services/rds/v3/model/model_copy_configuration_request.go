package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type CopyConfigurationRequest struct {
	ConfigId string `json:"config_id"`

	Body *ConfigurationCopyRequestBody `json:"body,omitempty"`
}

func (o CopyConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyConfigurationRequest struct{}"
	}

	return strings.Join([]string{"CopyConfigurationRequest", string(data)}, " ")
}
