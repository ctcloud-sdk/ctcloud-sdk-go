package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type UpdatePortResponse struct {
	Port           *Port `json:"port,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdatePortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePortResponse struct{}"
	}

	return strings.Join([]string{"UpdatePortResponse", string(data)}, " ")
}
