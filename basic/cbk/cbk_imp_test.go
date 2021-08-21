package cbk

import (
	"HelloGo/basic/cbk/util"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func init()  {
	log.SetFormatter(&log.TextFormatter{
		//DisableColors: true,
		FullTimestamp: true,
		TimestampFormat: time.StampMilli,
	})
}

const HOST_PREFIX = "http://www.abc.com"
const API_PREFIX = "/fake-api"

var REC_SIGN = make(chan struct{}, 1)
var REQ_CH = make(chan int, 1)

func TestCircuitBreakerImp(t *testing.T) {
	log.Infof("Test for cbk: %s", HOST_PREFIX+API_PREFIX)

	cbk := &CircuitBreakerImp{}
	cbk.apiMap = make(map[string]*apiSnapShop)
	// 控制时间窗口，10秒一轮, 重置api错误率
	cbk.roundInterval = util.ToDuration(10 * time.Second)
	// 熔断之后，3秒不出现错误再恢复
	cbk.recoverInterval = util.ToDuration(3 * time.Second)
	cbk.minCheck = 3
	cbk.cbkErrRate = 0.5

	// 持续失败
	go keepFailedReq()
	// 等待成功
	go waitForSuccess()
	//go reportStatus(cbk)
	StartJob(cbk, REQ_CH)
}

func waitForSuccess() {
	for {
		_ = <- REC_SIGN
		// mock for success
		log.Warnf("# Mock for success!")
		REQ_CH <- 1
		time.Sleep(1)
	}
}

func keepFailedReq() {
	for {
		// 每3秒发1次失败
		REQ_CH <- 0
		time.Sleep(3 * time.Second)
	}
}

func StartJob(cbk *CircuitBreakerImp, reqCh chan int) {
	for {
		time.Sleep(time.Second * 1)
		//tk := time.Tick(cbk.roundInterval * time.Second)
		select {
		//case <-tk:
			// reset
			//log.Warnf("With %v, Round finished...", cbk.roundInterval)
		case req := <-reqCh:
			if req == 1 {
				log.Warnf("Send success req id-%v!", req)
			}
			// do request
			ReqForTest(cbk, req)
		}
	}
}

func reportStatus(cbk *CircuitBreakerImp) {
	for {
		log.Warn("Report for cbk status...")
		for k, v := range cbk.apiMap {
			log.Warnf("Cbk map status: API: %v, isPaused: %v, errCount: %v,"+
				" total: %v, accessLast: %v, rountLast: %v", k, v.isPaused,
				v.errCount, v.totalCount, v.accessLast, v.roundLast)
		}
		time.Sleep(3 * time.Second)
	}
}

func ReqForTest(cbk *CircuitBreakerImp, req int) {
	// mock failed case
	mockAPI := API_PREFIX //+ strconv.Itoa(req)
	log.Infof("Ready to reqForTest: %s, req-id-%v", HOST_PREFIX+mockAPI, req)

	if !cbk.CanAccess(mockAPI, req, REC_SIGN) {
		log.Errorf("Api: %v is break, req-id-%v, wait for next round or success for one...", mockAPI, req)
		return
	} else {
		log.Infof("Continue ReqForTest: %s, req-id-%v", HOST_PREFIX+mockAPI, req)
	}

	if req == 0 {
		log.Infof("# Meet failed ReqForTest: %s", HOST_PREFIX+mockAPI)
		cbk.Failed(mockAPI)
	} else {
		log.Infof("# Meet success ReqForTest: %s", HOST_PREFIX+mockAPI)
		cbk.Succeed(mockAPI)
	}

}
