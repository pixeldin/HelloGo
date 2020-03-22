package sun

import "github.com/sirupsen/logrus"

func Process(args ...string)  {
	if len(args) < 2 {
		logrus.Error(Usage())
		return
	}

	//Do something
	logrus.Infof("Sun %s about %s miles.", args[0], args[1])

}

func Usage() string {
	return "Hint: sun <rise/down> <range>"
}
