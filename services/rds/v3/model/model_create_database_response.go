package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type CreateDatabaseResponse struct {
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseResponse struct{}"
	}

	return strings.Join([]string{"CreateDatabaseResponse", string(data)}, " ")
}
