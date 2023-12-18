package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type GetTaskDetailListRspJobsInstance struct {
	Id string `json:"id"`

	Name string `json:"name"`
}

func (o GetTaskDetailListRspJobsInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetTaskDetailListRspJobsInstance struct{}"
	}

	return strings.Join([]string{"GetTaskDetailListRspJobsInstance", string(data)}, " ")
}
