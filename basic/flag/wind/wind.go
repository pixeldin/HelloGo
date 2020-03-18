package wind

import "github.com/sirupsen/logrus"

//可变参数作为可选入参
func Process(args ...string)  {
	if len(args) < 2 {
		logrus.Error(Usage())
		return
	}

	//Do something
	logrus.Infof("Wind fly from %s, level %s.", args[0], args[1])
}

func Usage() string {
	return "Hint: wind <N/S/W/E> <level>"
}
