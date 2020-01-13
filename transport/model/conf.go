package model

import "github.com/obase/conf"

var (
	Xdebug                  = conf.OptiBool("ext.xdebug", false)
	PRODUCE_SIZE            = conf.OptiInt("ext.prodSize", 3)
	CONSUME_SIZE            = conf.OptiInt("ext.consSize", 5)
	QUEUE_SIZE            = conf.OptiInt("ext.queueSize", 30)
)
