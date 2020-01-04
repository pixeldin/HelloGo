package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	HttpStatus int `json:"httpstatus"`
	Data ResponseData `json:"data"`
	Messages string `json:"messages"`
	Status bool `json:"status"`
}

type ResponseData struct {
	Result []string `json:"result"`
	Flag string `json:"flag"`
	Mapping map[string]string `json:"map"`
}

func Train() {
	client := &http.Client{}
	url := "https://kyfw.12306.cn/otn/leftTicket/queryZ?leftTicketDTO.train_date=2020-01-26&leftTicketDTO.from_station=CBQ&leftTicketDTO.to_station=ZHQ&purpose_codes=ADULT"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	ua := `Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36`
	ck := `JSESSIONID=C8A503DFE418BA205CA278B8F8193DC2; BIGipServerotn=3671523594.64545.0000; RAIL_EXPIRATION=1577819290257; RAIL_DEVICEID=PGot9AAyvN23niVmtw8hPK-HPVfGQsJ37thsSvC6E9tpkJ8YnUwVN2DeDeVVryMUuwa6mW63c5Xf3lclLFOg14ssNWGbXzL7TfIYHqqmyfXkAwy4djL-PvDrX4qYz2TvUIio6MPPmRbix9oEhiWBvWcN9_iQS3Iw; BIGipServerpassport=854065418.50215.0000; route=6f50b51faa11b987e576cdb301e545c4; _jc_save_fromStation=%u6F6E%u6C55%2CCBQ; _jc_save_toStation=%u73E0%u6D77%2CZHQ; _jc_save_fromDate=2020-01-26; _jc_save_toDate=2019-12-28; _jc_save_wfdc_flag=dc`
	request.Header.Add("User-Agent", ua)
	request.Header.Add("Cookie", ck)

	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("Request err: ", err)
		return
	}
	s, _ := ioutil.ReadAll(response.Body) //把  body 内容读入字符串 s
	//var mp map[string]interface{}
	var rps Response
	err = json.Unmarshal(s, &rps)
	if err != nil {
		fmt.Println("Unmarshal err: ", err)
	}
	//fmt.Println(string(s))
	if response.StatusCode == 200 {
		fmt.Println("Request resonse data:")
		//fmt.Println(mp["data"])
		fmt.Println(rps)
	}
}
