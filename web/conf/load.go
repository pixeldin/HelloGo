package conf

import (
	"fmt"
	"github.com/obase/conf"
)


const (
	EXT = "pvp"
	REDIS_CONF = "redis"
)

func init()  {
	configs, ok := conf.GetSlice(REDIS_CONF)
	if !ok || len(configs) == 0 {
		return
	}
	for _, config := range configs {
		key1, ok := conf.ElemString(config, "key")
		if ok {
			fmt.Println(key1)
		}
		key2, ok := conf.ElemString(config, "address")
		if ok {
			fmt.Println(key2)
		}

	}
	pvp1 := conf.OptiString("pvp.pvp1", "pvp_match_result_")
	fmt.Println(pvp1)
	pvp2 := conf.OptiString("pvp.pvp2", "pvp_match_result_")
	fmt.Println(pvp2)
}

func LoadingSomething()  {
	fmt.Println("Loadconf...")
}