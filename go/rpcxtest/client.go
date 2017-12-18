package rpcxtest

import (
	"context"
	"time"

	"github.com/apex/log"
	"github.com/smallnest/rpcx"
)

type ArgsClient struct {
	A int
	B int
}

type ReplyClient struct {
	C int
}

func StartClient(a, b int) {
	log.Info("Starting Client...")
	s := &rpcx.DirectClientSelector{Network: "tcp", Address: "127.0.0.1:8972"}
	client := rpcx.NewClient(s)
	//client.Timeout = 5 * time.Second
	// client.DialTimeout = 10 * time.Second
	// client.Timeout = 10 * time.Second
	// client.ReadTimeout = 10 * time.Second
	// client.WriteTimeout = 10 * time.Second

	args := &ArgsClient{a, b}
	var reply ReplyClient
	err := client.Call(context.Background(), "Arith.Mul", args, &reply)
	if err != nil {
		log.Infof("error for Arith: %d*%d, %v", args.A, args.B, err)
	} else {
		log.Infof("Arith: %d*%d=%d", args.A, args.B, reply.C)
	}

	//client.Close()
	time.Sleep(6 * time.Second)
	err = client.Call(context.Background(), "Arith.Mul", args, &reply)
	if err != nil {
		log.Infof("error for Arith: %d*%d, %v", args.A, args.B, err)
	} else {
		log.Infof("Arith: %d*%d=%d", args.A, args.B, reply.C)
	}
	//client.Close()
	//log.Info("Client Closed")
}

func StartClient1() {
	log.Info("Starting Client1...")
	s := &rpcx.DirectClientSelector{Network: "tcp", Address: "127.0.0.1:8972"}
	client := rpcx.NewClient(s)
	//client.Timeout = 5 * time.Second
	// client.DialTimeout = 10 * time.Second
	// client.Timeout = 10 * time.Second
	// client.ReadTimeout = 10 * time.Second
	// client.WriteTimeout = 10 * time.Second

	args := &ArgsClient{8, 9}
	var reply ReplyClient
	err := client.Call(context.Background(), "Arith.Mul", args, &reply)
	if err != nil {
		log.Infof("error for Arith: %d*%d, %v", args.A, args.B, err)
	} else {
		log.Infof("Arith: %d*%d=%d", args.A, args.B, reply.C)
	}

	// //client.Close()
	// time.Sleep(6 * time.Second)
	// err = client.Call(context.Background(), "Arith.Mul", args, &reply)
	// if err != nil {
	// 	log.Infof("error for Arith: %d*%d, %v", args.A, args.B, err)
	// } else {
	// 	log.Infof("Arith: %d*%d=%d", args.A, args.B, reply.C)
	// }
	//client.Close()
	//log.Info("Client Closed")
}

func StartLoopClient() {
	log.Info("Starting Loop Client...")
	s := &rpcx.DirectClientSelector{Network: "tcp", Address: "127.0.0.1:8972"}
	client := rpcx.NewClient(s)

	var reply ReplyClient
	num := 1
	for {
		args := &ArgsClient{num, num}
		num++
		err := client.Call(context.Background(), "Arith.Mul", args, &reply)
		if err != nil {
			log.Infof("error for Arith: %d*%d, %v", args.A, args.B, err)
		} else {
			log.Infof("Arith: %d*%d=%d", args.A, args.B, reply.C)
		}
		time.Sleep(3 * time.Second)
	}

}
