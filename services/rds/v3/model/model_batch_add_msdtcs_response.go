package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type BatchAddMsdtcsResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchAddMsdtcsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddMsdtcsResponse struct{}"
	}

	return strings.Join([]string{"BatchAddMsdtcsResponse", string(data)}, " ")
}
