package connect

import (
	"bufio"
	"context"
	"errors"
	"math"
	"net"
	"runtime"
	"strconv"
	"time"

	"DailysServer/connect/conf"
	"DailysServer/pkg/timer"
	"DailysServer/proto"

	"github.com/freezeChen/studio-library/zlog"
)

func InitTCP(s *Server, c *conf.TCPConfig) (err error) {

	var (
		addr *net.TCPAddr
	)

	addr, err = net.ResolveTCPAddr("tcp", c.Port)
	if err != nil {
		zlog.Errorf("net.ResolveTCPAddr error(%s)", err)
		return
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		zlog.Errorf("net.listenTCP error(%s)", err)
		return
	}

	for i := 0; i < runtime.NumCPU(); i++ {
		go s.acceptTCP(c, listener)
	}

	return
}

func (server *Server) acceptTCP(c *conf.TCPConfig, listen *net.TCPListener) {
	var (
		conn *net.TCPConn
		err  error
		r    int
	)

	for {
		//等待连接
		if conn, err = listen.AcceptTCP(); err != nil {
			zlog.Errorf("listen.acceptTCP error(%s)", err)
			return
		}
		if err := conn.SetKeepAlive(c.TCPKeepalive); err != nil {
			zlog.Errorf("conn setKeepAlive error(%s)", err)
			return
		}

		if err := conn.SetReadBuffer(c.TCPReadBuffer); err != nil {
			zlog.Errorf("conn setReadBuffer error(%s)", err)
			return
		}

		if err := conn.SetWriteBuffer(c.TCPWriteBuffer); err != nil {
			zlog.Errorf("conn setWriterBuffer error(%s)", err)
			return
		}

		go server.serverTCP(c, conn, r)
		if r++; r == math.MaxInt32 {
			r = 0
		}
	}
}

func (server *Server) serverTCP(c *conf.TCPConfig, conn *net.TCPConn, r int) {
	var (
		ch  *Channel
		msg *proto.Proto
		err error
		//ctx context.Context
		//lastHb    = time.Now()
		tim       = server.round.Timer(r)
		timerData *timer.TimerData
	)

	ch = NewChannel()
	ch.Reader = *bufio.NewReader(conn)
	ch.Writer = *bufio.NewWriter(conn)

	ch.Conn = conn

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timerData = tim.Add(_HandshakeTimeout, func() {
		zlog.Errorf("shakeHands timeout")
		conn.Close()
	})

	if msg, err = ch.Ring.Set(); err == nil {
		if ch.Id, err = server.AuthTCP(ctx, msg, ch); err == nil {
			server.Bucket.Online(ch.Id, ch)
			zlog.Debugf("tcp connect id:%d proto: %+v", ch.Id, msg)
		}

		if err != nil {
			ch.Close()
			conn.Close()
			tim.Del(timerData)
			zlog.Errorf("id:%d handshake error(%v)", ch.Id, err)
			return
		}

		timerData.Key = ch.Id
		tim.Set(timerData, time.Duration(10*time.Second))
		go server.dispatchTCP(conn, &ch.Writer, ch)

		for {
			if msg, err = ch.Ring.Set(); err != nil {
				break
			}

			if err := msg.ReadTCP(&ch.Reader); err != nil {
				break
			}

			if msg.Opr == proto.OpHeartbeat {
				zlog.Info("heartbeat")
				tim.Set(timerData, _HeartBeat)
				msg.Opr = proto.OpHeartbeatReply
				msg.Body = nil

			} else {
				if err := server.Operate(ctx, msg); err != nil {
					break
				}
			}

			ch.Ring.SetAdv()
			ch.Signal()

		}
	} else {
		zlog.Errorf("set error(%s)", err)
	}

	tim.Del(timerData)
	conn.Close()
	ch.Close()
	server.Bucket.Offline(ch.Id)

}

func (server *Server) dispatchTCP(conn *net.TCPConn, wr *bufio.Writer, ch *Channel) {
	var (
		err error
	)

	for {
		p := ch.Ready()
		switch p {
		case proto.ProtoFinish:
			goto failed
		case proto.ProtoReady:
			if p, err = ch.Ring.Get(); err != nil {
				break
			}

			if p.Opr == proto.OpHeartbeatReply {

			}

			if err := p.WriteTCP(wr); err != nil {
				goto failed
			}

			p.Body = nil
			ch.Ring.GetAdv()
		default:

		}

		if err := wr.Flush(); err != nil {
			goto failed
		}
	}

failed:
	conn.Close()
}

func (server *Server) AuthTCP(ctx context.Context, msg *proto.Proto, ch *Channel) (id int64, err error) {
	zlog.Debug("wait for auth in tcp")
	if err := msg.ReadTCP(&ch.Reader); err != nil {
		return 0, err
	}

	if msg.Opr != proto.OpAuth {
		err = errors.New("authTCP op is error")
		return
	}

	id, err = strconv.ParseInt(string(msg.Body), 10, 64)
	if err != nil {
		return
	}

	msg.Opr = proto.OpAuthReply
	msg.Body = nil

	if err := msg.WriteTCP(&ch.Writer); err != nil {
		return 0, err
	}

	if err = server.Connect(ctx, conf.Conf.RpcServer.Id, id); err != nil {
		return 0, err
	}

	return
}
