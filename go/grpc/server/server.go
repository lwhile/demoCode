package server

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	context "golang.org/x/net/context"

	"demo/grpc/message"
	pb "demo/grpc/message"
	"strconv"

	"git.gzsunrun.cn/sunrunlib/log"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// Start will start a server
func Start() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Error(err)
		return
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEchoServer(grpcServer, &server{})
	go func() {
		time.Sleep(1 * time.Second)
		log.Infoln("Start gRPC server success")
		go startTrace()
	}()

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}

}

func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
	go http.ListenAndServe(":50051", nil)
	grpclog.Println("Trace listen on 50051")
}

type server struct{}

// Echo :
func (s *server) Echo(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	//name, _ := ctx.Value("name").(string)
	message := in.GetMessage()
	return &pb.Response{Message: message}, nil
}

func (s *server) ClientStream(stream message.Echo_ClientStreamServer) error {
	var reqs []*message.Request
	for {
		req, err := stream.Recv()
		log.Infof("Server received data %q from stream client\n", req)
		if err == io.EOF {
			var sum int
			for _, ele := range reqs {
				n, _ := strconv.Atoi(ele.Message)
				sum += n
			}
			return stream.SendAndClose(&message.Response{Message: strconv.Itoa(sum)})
		}
		if err != nil {
			return err
		}

		reqs = append(reqs, req)
	}
}

func (s *server) ServerStream(r *message.Request, stream message.Echo_ServerStreamServer) error {
	num, err := strconv.Atoi(r.Message)
	if err != nil {
		return err
	}

	primes := primeNumber(num)

	for _, ele := range primes {
		if err := stream.Send(&message.Response{Message: strconv.Itoa(ele)}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 500)
	}
	return nil
}

func (s *server) BilateralStream(stream message.Echo_BilateralStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Infof("Server received message from client: %s\n", req.Message)

		resp := message.Response{Message: fmt.Sprintf("%d", time.Now().Unix())}

		if err = stream.Send(&resp); err != nil {
			return nil
		}
	}
}
