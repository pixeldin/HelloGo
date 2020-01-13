package worker

import (
	"github.com/sirupsen/logrus"
	"time"
)

func Consume(CNum int, msg chan []byte)  {
	for value := range msg {
		//Add time costing
		time.Sleep(2 * time.Second)
		logrus.Infof("# Consumer CNum.%d, take cake with value: %s.", CNum, string(value))
	}
}
