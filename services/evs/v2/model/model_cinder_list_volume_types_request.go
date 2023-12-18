package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type CinderListVolumeTypesRequest struct {
}

func (o CinderListVolumeTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListVolumeTypesRequest struct{}"
	}

	return strings.Join([]string{"CinderListVolumeTypesRequest", string(data)}, " ")
}
