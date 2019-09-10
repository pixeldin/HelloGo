package conf

import (
	"fmt"
	"github.com/obase/conf"
)

const (
	EXT      = "pvp"
	CONF_KEY = "redis"
)


func init()  {
	configs, ok := conf.GetSlice(CONF_KEY)
	if !ok || len(configs) == 0 {
		return
	}
	for _, config := range configs {
		addr, ok := conf.ElemString(config, "address")
		if ok {
			fmt.Println(addr)
		}
		pwd, ok := conf.ElemString(config, "password")
		if ok {
			fmt.Println(pwd)
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