package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type SetReadOnlySwitchResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetReadOnlySwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetReadOnlySwitchResponse struct{}"
	}

	return strings.Join([]string{"SetReadOnlySwitchResponse", string(data)}, " ")
}
