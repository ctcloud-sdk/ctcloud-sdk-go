package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type NovaAttachInterfaceRequestBody struct {
	InterfaceAttachment *NovaAttachInterfaceOption `json:"interfaceAttachment"`
}

func (o NovaAttachInterfaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAttachInterfaceRequestBody struct{}"
	}

	return strings.Join([]string{"NovaAttachInterfaceRequestBody", string(data)}, " ")
}
