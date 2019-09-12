package connect

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"time"

	"DailysServer/pkg/timer"
	"DailysServer/proto"

	"github.com/freezeChen/studio-library/zlog"

	"github.com/gorilla/websocket"
)

var rn int

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
}

func InitWebSocket(srv *Server, addr string) {
	http.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		upGrade(srv, w, r)
	})
	go func() {
		if err := http.ListenAndServe(addr, nil); err != nil {
			panic(err)
			return
		}
	}()

}

func upGrade(srv *Server, w http.ResponseWriter, r *http.Request) {
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		zlog.Errorf(err.Error())
		return
	}
	go ServeWebSocket(srv, conn, rn)
	if rn++; rn == math.MaxInt32 {
		rn = 0
	}
}

func ServeWebSocket(srv *Server, ws *websocket.Conn, rn int) {
	var (
		err      error
		ch       = NewChannel()
		key      string
		msg      *proto.Proto
		tim      = srv.round.Timer(rn)
		timeData *timer.TimerData
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeData = tim.Add(_HandshakeTimeout, func() {
		zlog.Debug("timeout")
		ws.Close()
	})

	if msg, err = ch.Ring.Set(); err == nil {

		if key, err = srv.AuthWebSocket(ctx, ws, msg, ch); err == nil {
			srv.bucket.Online(key, ch)
			ch.Id = key
			msg.Opr = proto.OpAuthReply
			zlog.Debug("auth:" + key)
		}
	}

	if err != nil {
		zlog.Error(err.Error())
		ch.Close()
		ws.Close()
		tim.Del(timeData)
		return
	}

	timeData.Key = ch.Id
	tim.Set(timeData, 60*time.Second)

	go srv.dispatchWebsocket(ws, ch)

	for {
		if msg, err = ch.Ring.Set(); err != nil {
			break
		}
		if err = msg.ReadWebSocket(ws); err != nil {
			zlog.Error(err.Error())
			break
		}
		if msg.Opr == proto.OpHeartbeat {
			tim.Set(timeData, _HeartBeat)
			msg.Opr = proto.OpHeartbeatReply
			msg.Body = nil
			zlog.Debug("ping")
		}

		ch.Ring.SetAdv()
		ch.Signal()
	}
	zlog.Debug("close2" + err.Error())
	ch.Close()
	ws.Close()
	tim.Del(timeData)
	srv.bucket.Offline(key)

}

func (s *Server) dispatchWebsocket(ws *websocket.Conn, ch *Channel) {

	var err error
	for {
		msg := ch.Ready()
		switch msg {
		case proto.ProtoFinish:
			zlog.Info("websocket finish")
			goto field1
		case proto.ProtoReady:
			if msg, err = ch.Ring.Get(); err != nil {
				goto field1
			}

			if err = msg.WriteWebSocket(ws); err != nil {
				goto field1
			}
			msg.Body = nil
			ch.Ring.GetAdv()
		default:

			zlog.Debug("default")
			if err := msg.WriteWebSocket(ws); err != nil {
				goto field1
			}
		}
	}

field1:
	zlog.Errorf("dispatchTCP error(%s)", err)
	ws.Close()
	return

}

func (s *Server) AuthWebSocket(ctx context.Context, ws *websocket.Conn, msg *proto.Proto, ch *Channel) (key string, err error) {
	fmt.Println("auth")
	err = msg.ReadWebSocket(ws)
	if err != nil {
		return
	}

	/*	key = string(msg.Body)

		zlog.Debug("auth:" + key)

		msg.Opr = OpAuthReply
		if err = msg.WriteWebSocket(ws); err != nil {
			return
		}

		var uid int64

		if uid, err = strconv.ParseInt(key, 10, 64); err != nil {
			return
		}


		return
	*/
	//校验token

	return
}
