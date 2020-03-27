package model

import (
	"github.com/obase/conf"
	"time"
)

var (
	PROCESS_TICK_TIME = conf.OptiDuration("ext.processTick", 3*time.Second)
)
