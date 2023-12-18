package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type DeletePostgresqlExtensionResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePostgresqlExtensionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePostgresqlExtensionResponse struct{}"
	}

	return strings.Join([]string{"DeletePostgresqlExtensionResponse", string(data)}, " ")
}
