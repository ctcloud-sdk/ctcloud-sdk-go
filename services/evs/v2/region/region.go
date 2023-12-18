package region

import (
	"fmt"
	"github.com/ctcloud-sdk/ctcloud-sdk-go/core/region"
)

var (
	RU_MOSCOW_1 = region.NewRegion("ru-moscow-1",
		"https://evs.ru-moscow-1.hc.sbercloud.ru")
)

var staticFields = map[string]*region.Region{
	"ru-moscow-1": RU_MOSCOW_1,
}

var provider = region.DefaultProviderChain("EVS")

func ValueOf(regionId string) *region.Region {
	if regionId == "" {
		panic("unexpected empty parameter: regionId")
	}

	reg := provider.GetRegion(regionId)
	if reg != nil {
		return reg
	}

	if _, ok := staticFields[regionId]; ok {
		return staticFields[regionId]
	}
	panic(fmt.Sprintf("unexpected regionId: %s", regionId))
}
