package sockio

import (
	"log"
	"net/http"
	"time"

	"github.com/google/wire"

	"github.com/heromicro/omgind/pkg/global"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"

	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

// ProviderSet 注入
var ProviderSet = wire.NewSet(New)

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func New() (*socketio.Server, func(), error) {
	cfg := global.CFG

	//log.Println(" --== cfg:", cfg)
	//server := socketio.NewServer(nil)
	transports := []transport.Transport{
		&websocket.Transport{
			CheckOrigin: allowOriginFunc,
		},
	}
	server := socketio.NewServer(&engineio.Options{
		Transports:   transports,
		PingInterval: time.Second * 20,
	})

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println(" --== connected:", s.ID())
		log.Println(" ------- ==== id: ", s.ID(), " connected at: ", time.Now().UnixNano())

		s.Emit("ok", "welcome")
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println(" ====000-- OnError error:", e)
	})

	server.OnEvent("/", "start", func(s socketio.Conn, msg string) {
		log.Println(" ------ ==== msg: ", msg)

	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println(" ===== 99999 = closed ", reason)
	})

	server.Adapter(&socketio.RedisAdapterOptions{
		Addr:   cfg.Redis.DSN(),
		Prefix: "socket.io",
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	cleanFunc := func() {
		server.Close()
		log.Println(" clean in sockio ")
	}
	return server, cleanFunc, nil
}
