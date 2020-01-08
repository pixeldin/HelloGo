package model

import "github.com/obase/conf"

var (
	Xdebug                  = conf.OptiBool("ext.xdebug", false)
	PRODUCE_SIZE            = conf.OptiInt("ext.prodSize", 50)
	CONSUME_SIZE            = conf.OptiInt("ext.consSize", 2)
	QUEUE_SIZE            = conf.OptiInt("ext.queueSize", 300)
)
