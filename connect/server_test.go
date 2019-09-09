/*
   @Time : 2019-07-25 15:26
   @Author : frozenChen
   @File : server_test
   @Software: KingMaxWMS_APP_API
*/
package server

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"24on/account_api/conf"
	"gitee.com/bethink1501/24on-library/tokens"
	"gitee.com/bethink1501/24on-library/zlog"
)

func TestWebSocket(t *testing.T) {
	conf.InitConfig()
	zlog.InitLogger(conf.GetEnv().Log)

	s := NewServer(tokens.Init(tokens.KeyName("FUserID"), tokens.ValueKey([]string{"id"})))

	timer := time.NewTimer(1 * time.Second)

	go func() {
		for {
			<-timer.C
			timer.Reset(5*time.Second)
			fmt.Println("auth send")
			s.batchPush([]byte("count"), []string{"5"})
		}
	}()
	//time.Timer{
	//	C: nil,
	//}

	http.HandleFunc("/web", func(writer http.ResponseWriter, request *http.Request) {
		token := request.Header["Authorization"]
		//jwt.re
		fmt.Println(token)

		InitWebSocket(s, writer, request)
	})

	http.HandleFunc("/send", func(writer http.ResponseWriter, request *http.Request) {
		//request
		request.ParseForm()

		strings := request.Form["msg"][0]
		id := request.Form["id"]
		fmt.Println("msg:"+strings+",id:", id)

		s.batchPush([]byte(strings), id)
	})

	http.ListenAndServe(":8888", nil)

}
