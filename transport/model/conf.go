package model

import "github.com/obase/conf"

var (
	Xdebug                  = conf.OptiBool("ext.xdebug", false)
	PRODUCE_SIZE            = conf.OptiInt("ext.prodSize", 5)
	CONSUME_SIZE            = conf.OptiInt("ext.consSize", 5)
)
