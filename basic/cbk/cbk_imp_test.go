package cbk

import (
	"HelloGo/basic/cbk/util"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

const HOST_PREFIX = "http://www.abc.com"
const API_PREFIX = "/fake-api"

var REC_SIGN = make(chan struct{}, 1)

func TestCircuitBreakerImp(t *testing.T) {
	log.Infof("Test for cbk: %s", HOST_PREFIX+API_PREFIX)

	cbk := &CircuitBreakerImp{}
	cbk.apiMap = make(map[string]*apiSnapShop)
	// 10秒一轮
	cbk.roundInterval = util.ToDuration(10 * time.Second)
	// 3秒不出现错误恢复
	cbk.recoverInterval = util.ToDuration(2 * time.Second)
	cbk.minCheck = 5
	cbk.cbkErrRate = 0.5

	reqCh := make(chan int, 1)
	// 持续失败
	go keepFailedReq(reqCh)
	// 等待成功
	go waitForSuccess(reqCh)
	//go reportStatus(cbk)
	StartJob(cbk, reqCh)
}

func waitForSuccess(reqCh chan int) {
	for {
		_ = <- REC_SIGN
		// mock for success
		log.Warnf("# Mock for success!")
		reqCh <- 1
		time.Sleep(1)
	}
}

func keepFailedReq(recChan chan int) {
	for {
		// 每1秒发2次失败
		recChan <- 0
		time.Sleep(500 * time.Millisecond)
	}
}

func StartJob(cbk *CircuitBreakerImp, reqCh chan int) {
	for {
		//time.Sleep(time.Second * 1)
		tk := time.Tick(cbk.roundInterval * time.Second)
		select {
		case <-tk:
			// reset
			log.Warnf("With %v, Round finished...", cbk.roundInterval)
		case req := <-reqCh:
			// req.do
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
	log.Infof("Ready to reqForTest: %s", HOST_PREFIX+mockAPI)

	if !cbk.CanAccess(mockAPI, REC_SIGN) {
		log.Errorf("Api: %v is break, wait for next round or success for one...", mockAPI)
		return
	} else {
		log.Infof("Continue ReqForTest: %s", HOST_PREFIX+mockAPI)
	}

	if req == 0 {
		cbk.Failed(mockAPI)
	} else {
		log.Infof("# Meet Success ReqForTest: %s", HOST_PREFIX+mockAPI)
		cbk.Succeed(mockAPI)
	}

}
