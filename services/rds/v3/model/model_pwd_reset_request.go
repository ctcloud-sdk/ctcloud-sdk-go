package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type PwdResetRequest struct {
	DbUserPwd string `json:"db_user_pwd"`
}

func (o PwdResetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PwdResetRequest struct{}"
	}

	return strings.Join([]string{"PwdResetRequest", string(data)}, " ")
}
