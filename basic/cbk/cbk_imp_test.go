package cbk

import (
	"HelloGo/basic/cbk/util"
	log "github.com/sirupsen/logrus"
	"strconv"
	"testing"
	"time"
)

const HOST_PREFIX = "http://www.abc.com"
const API_PREFIX = "/fake-api-"

func TestCircuitBreakerImp(t *testing.T) {
	log.Infof("Test for cbk: %s", HOST_PREFIX+API_PREFIX)

	cbk := &CircuitBreakerImp{}
	cbk.apiMap = make(map[string]*apiSnapShop)
	// 10秒一轮
	cbk.roundInterval = util.ToDuration(10 * time.Second)
	// 3秒不出现错误恢复
	cbk.recoverInterval = util.ToDuration(3 * time.Second)
	cbk.minCheck = 10
	cbk.cbkErrRate = 0.5

	reqCh := make(chan int, 1)
	recSign := make(chan struct{}, 1)
	go sendReq(reqCh, recSign)
	go reportStatus(cbk)
	StartJob(cbk, reqCh, recSign)
}

func sendReq(recChan chan int, recoverSign chan struct{}) {
	for {
		//tk := time.Tick(100 * time.Millisecond)
		select {
		// 熔断半关闭则尝试一次成功
		//case <-tk:
		case <-recoverSign:
			recChan <- 1
		default:
			// 每1秒发10次失败
			recChan <- 0
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func StartJob(cbk *CircuitBreakerImp, reqCh chan int, recSign chan struct{}) {
	for {
		//time.Sleep(time.Second * 1)
		tk := time.Tick(cbk.roundInterval * time.Second)
		select {
		case <-tk:
			// reset
			log.Warnf("With %v, Round finished...", cbk.roundInterval)
		case req := <-reqCh:
			// req.do
			ReqForTest(cbk, req, recSign)
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

func ReqForTest(cbk *CircuitBreakerImp, req int, recSign chan struct{}) {
	// mock failed case
	mockAPI := API_PREFIX + strconv.Itoa(req)
	log.Infof("Ready to reqForTest: %s", HOST_PREFIX+mockAPI)

	if !cbk.CanAccess(mockAPI) {
		log.Errorf("Api: %v is break, wait for next round or success for one...", mockAPI)
		return
	} else {
		log.Infof("Continue ReqForTest: %s", HOST_PREFIX+mockAPI)
		// 通知外部成功一次
		recSign <- struct{}{}
	}

	if req == 0 {
		cbk.Failed(mockAPI)
	} else {
		log.Infof("# Meet Success ReqForTest: %s", HOST_PREFIX+mockAPI)
		cbk.Succeed(mockAPI)
	}

}
