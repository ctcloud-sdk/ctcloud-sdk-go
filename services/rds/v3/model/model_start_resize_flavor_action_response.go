package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type StartResizeFlavorActionResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartResizeFlavorActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartResizeFlavorActionResponse struct{}"
	}

	return strings.Join([]string{"StartResizeFlavorActionResponse", string(data)}, " ")
}
