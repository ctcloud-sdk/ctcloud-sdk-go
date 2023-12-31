package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type BatchStopServersRequestBody struct {
	OsStop *BatchStopServersOption `json:"os-stop"`
}

func (o BatchStopServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopServersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchStopServersRequestBody", string(data)}, " ")
}
