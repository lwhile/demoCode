package client

import (
	"fmt"
	"io"
	"time"

	"git.gzsunrun.cn/sunrunlib/log"
	"google.golang.org/grpc"

	pb "demo/grpc/message"

	"demo/grpc/message"
	"strconv"

	"golang.org/x/net/context"
)

func StartEcho() {
	log.Infoln("starting client...")
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)

	r, err := client.Echo(context.WithValue(context.Background(), "Message", "lwh"), &pb.Request{Message: "lwh"})
	if err != nil {
		log.Error(err)
		return
	}

	log.Infoln(r.GetMessage())
}

func StartClientStream() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return
	}
	client := pb.NewEchoClient(conn)

	stream, err := client.ClientStream(context.Background())
	if err != nil {
		log.Error(err)
	}

	// sent numbers from 1 to 100
	for i := 0; i < 100; i++ {
		if err := stream.Send(&message.Request{Message: strconv.Itoa(i)}); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	log.Info(reply.Message)
}

func StartServerStream() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return
	}
	client := pb.NewEchoClient(conn)
	stream, err := client.ServerStream(context.Background(), &message.Request{Message: "10"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Info(resp)
	}
}

func BilateralStream() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		return
	}
	client := pb.NewEchoClient(conn)

	stream, err := client.BilateralStream(context.Background())
	if err != nil {
		log.Error(err)
		return
	}

	if err := stream.Send(&message.Request{Message: fmt.Sprintf("%d", time.Now().Unix())}); err != nil {
		return
	}

	for {
		resp, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				log.Error(err)
			}
			return
		}

		log.Infof("Client received message from server: %s\n", resp.Message)

		req := message.Request{Message: fmt.Sprintf("%d", time.Now().Unix())}
		if err = stream.Send(&req); err != nil {
			log.Error(err)
			return
		}
		time.Sleep(time.Second)
	}
}
