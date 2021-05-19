package consumeQ

import (
	"testing"
	"time"
)

func TestMqBackupToMongo(t *testing.T) {
	MqBackupToMongo()
	for {
		time.Sleep(1)
	}
}
