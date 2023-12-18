package model

import (
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/utils"

	"strings"
)

type NovaListKeypairsResponse struct {
	Keypairs       *[]NovaListKeypairsResult `json:"keypairs,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o NovaListKeypairsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListKeypairsResponse struct{}"
	}

	return strings.Join([]string{"NovaListKeypairsResponse", string(data)}, " ")
}
