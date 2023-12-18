package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type BinlogClearPolicyRequestBody struct {
	BinlogRetentionHours int64 `json:"binlog_retention_hours"`
}

func (o BinlogClearPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BinlogClearPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"BinlogClearPolicyRequestBody", string(data)}, " ")
}
