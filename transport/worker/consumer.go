package worker

import (
	"github.com/sirupsen/logrus"
	"time"
)

func Consume(CNum int, msg chan []byte)  {
	for value := range msg {
		logrus.Infof("# Consumer CNum.%d, take cake, value: %s.", CNum, string(value))
		//Add time costing
		time.Sleep(2000)
	}
}
