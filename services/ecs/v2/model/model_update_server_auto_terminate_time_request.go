package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type UpdateServerAutoTerminateTimeRequest struct {
	ServerId string `json:"server_id"`

	Body *UpdateServerAutoTerminateTimeRequestBody `json:"body,omitempty"`
}

func (o UpdateServerAutoTerminateTimeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerAutoTerminateTimeRequest struct{}"
	}

	return strings.Join([]string{"UpdateServerAutoTerminateTimeRequest", string(data)}, " ")
}