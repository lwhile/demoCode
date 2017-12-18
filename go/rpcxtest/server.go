package rpcxtest

import (
	"context"
	"time"

	"github.com/apex/log"
	"github.com/smallnest/rpcx"
)

type ArgsServer struct {
	A int
	B int
}

type ReplyServer struct {
	C int
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *ArgsServer, reply *ReplyServer) error {
	reply.C = args.A * args.B
	//time.Sleep(time.Minute)
	return nil
}

func StartServer() {
	log.Info("Starting Server")
	server := rpcx.NewServer()
	// server.Timeout = 5 * time.Second
	// server.ReadTimeout = 5 * time.Second
	// server.WriteTimeout = 5 * time.Second
	go func() {
		time.Sleep(2 * time.Second)
		server.Close()
	}()
	server.RegisterName("Arith", new(Arith))
	server.Serve("tcp", "127.0.0.1:8972")
}

// func StartServer1() {
// 	log.Info("Starting Server1")
// 	server := rpcx.NewServer()
// 	//server.Timeout = 5 * time.Second
// 	//server.ReadTimeout = 5 * time.Second
// 	//server.WriteTimeout = 5 * time.Second

// 	server.RegisterName("Arith", new(Arith))
// 	server.Serve("tcp", "127.0.0.1:8972")
// }
