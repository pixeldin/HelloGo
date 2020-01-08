package worker

import (
	"github.com/sirupsen/logrus"
)

func Consume(CNum int, msg chan []byte)  {
	for value := range msg {
		logrus.Infof("# Consumer CNum.%d, take cake, value: %s.", CNum, string(value))
	}
}
